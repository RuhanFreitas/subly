package main

import (
	"log"
	"net/http"
	"subly/internal/config"
	"subly/internal/database"
	"subly/internal/handler"
	midleware "subly/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration")
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	var router *gin.Engine = gin.Default()

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Health": "The API is healthy!",
		})
	})

	router.POST("/auth/register", handler.Register(pool, cfg))
	router.POST("auth/login", handler.Login(pool, cfg))

	router.GET("/user/:id", midleware.AuthMiddleware(cfg), handler.GetUserByID(pool))
	router.PATCH("/user/:id", midleware.AuthMiddleware(cfg), handler.UpdateUserByID(pool))
	router.DELETE("/user/:id", midleware.AuthMiddleware(cfg), handler.DeleteUser(pool))

	router.POST("/subscription", midleware.AuthMiddleware(cfg), handler.CreateSubscription(pool))
	router.GET("/subscription", midleware.AuthMiddleware(cfg), handler.GetAllSubscriptions(pool))
	router.GET("/subscription/:id", midleware.AuthMiddleware(cfg), handler.GetSubscriptionByID(pool))
	router.PATCH("/subscription/:id", midleware.AuthMiddleware(cfg), handler.UpdateSubscription(pool))
	router.DELETE("/subscription/:id", midleware.AuthMiddleware(cfg), handler.DeleteSubscription(pool))

	router.Run(":" + cfg.Port)
}
