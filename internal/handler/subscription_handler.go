package handler

import (
	"net/http"
	"strconv"
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
	SubscriptionRenew string    `json:"subscription_renew" binding:"required"`
}

type UpdateSubscriptionInput struct {
	Name              *string    `json:"name"`
	Price             *float64   `json:"price"`
	StartingDate      *time.Time `json:"starting_date"`
	PaymentDate       *time.Time `json:"payment_date"`
	SubscriptionRenew *string    `json:"subscription_renew"`
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried" + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"susbcription": subscription})
	}
}

func GetAllSubscriptions(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
			return
		}

		var subscriptions *[]model.Subscription

		subscriptions, err = repository.GetAllSubscriptions(pool, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
	}
}

func GetSubscriptionByID(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
			return
		}

		var subscription *model.Subscription

		subscription, err = repository.GetSubscriptionByID(pool, id)
		if err != nil {
			if err.Error() == "no rows in result set" {
				c.JSON(http.StatusNotFound, gin.H{"error": "No subscription found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"subscription": subscription})
	}
}

func UpdateSubscription(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
			return
		}

		var updateInput UpdateSubscriptionInput

		if err := c.ShouldBindJSON(&updateInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		var subscription *model.Subscription
		subscription, err = repository.GetSubscriptionByID(pool, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription not found"})
			return
		}

		name := subscription.Name
		if updateInput.Name != nil {
			name = *updateInput.Name
		}

		price := subscription.Price
		if updateInput.Price != nil {
			price = *updateInput.Price
		}

		startingDate := subscription.StartingDate
		if updateInput.StartingDate != nil {
			startingDate = *updateInput.StartingDate
		}

		paymentDate := subscription.PaymentDate
		if updateInput.PaymentDate != nil {
			paymentDate = *updateInput.PaymentDate
		}

		subscriptionRenew := subscription.SubscriptionRenew
		if updateInput.SubscriptionRenew != nil {
			subscriptionRenew = *updateInput.SubscriptionRenew
		}

		subscription, err = repository.UpdateSubscriptionByID(pool, id, name, price, startingDate, paymentDate, subscriptionRenew)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"subscription": subscription})
	}
}

func DeleteSubscription(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
			return
		}

		err = repository.DeleteSubscription(pool, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Subscription deleted successfully"})
	}
}
