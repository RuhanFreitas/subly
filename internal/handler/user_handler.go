package handler

import (
	"net/http"
	"strconv"
	"subly/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserInput struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
}

func UpdateUserByID(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access data"})
			return
		}

		var input UpdateUserInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		if input.FirstName == nil && input.LastName == nil && input.Email == nil && input.Password == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "At least one field must be provided"})
			return
		}

		existingUser, err := repository.GetUserByID(pool, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		firstName := existingUser.FirstName
		if input.FirstName != nil {
			firstName = *input.FirstName
		}

		lastName := existingUser.LastName
		if input.LastName != nil {
			lastName = *input.LastName
		}

		email := existingUser.Email
		if input.Email != nil {
			email = *input.Email
		}

		password := existingUser.Password
		if input.Password != nil {
			password = *input.Password
		} else {
			var hashedPassword []byte
			hashedPassword, err = bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
				return
			}
			password = string(hashedPassword)
		}

		updatedUser, err := repository.UpdateUserByID(pool, id, firstName, lastName, email, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": updatedUser})
	}
}

func DeleteUser(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid access data"})
			return
		}

		err = repository.DeleteUser(pool, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error occuried"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
	}
}
