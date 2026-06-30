package main

import (
	"log"
	"net/http"
	"subly/internal/config"
	"subly/internal/database"
	"subly/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// migrate create -ext sql -dir migrations -seq create_users_table
// migrate -path migrations -database $env:DATABASE_URL up
// migrate -path migrations -database "postgresql://appuser:apppassword@localhost:15432/appdb?sslmode=disable" up

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

	// Auth
	router.POST("/auth/register", handler.Register(pool, cfg))

	// User

	// Subscription

	router.Run(":8081")
}
