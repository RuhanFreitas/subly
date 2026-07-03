package handler

import (
	"net/http"
	"subly/internal/model"
	"subly/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// The user ID is currently coming from the params, but we need to change that later
// using the JWT and adding the user id in the request body so they don't need to parse ids

type SubscriptionInput struct {
	Name              string    `json:"name" binding:"required"`
	Price             float64   `json:"price" binding:"required"`
	StartingDate      time.Time `json:"starting_date" binding:"required"`
	PaymentDate       time.Time `json:"payment_date" binding:"required"`
	SubscriptionRenew time.Time `json:"subscription_renew" binding:"required"`
}

func CreateSubscription(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var input *SubscriptionInput

		err := c.ShouldBindJSON(&input)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		var subscription *model.Subscription = &model.Subscription{
			UserID:            id,
			Name:              input.Name,
			Price:             input.Price,
			StartingDate:      input.StartingDate,
			PaymentDate:       input.PaymentDate,
			SubscriptionRenew: input.SubscriptionRenew,
		}

		subscription, err = repository.CreateSubscription(pool, subscription)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"susbcription": subscription})
	}
}
