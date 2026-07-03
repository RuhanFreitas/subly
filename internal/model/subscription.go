package model

import "time"

type Subscription struct {
	ID                string    `json:"id" db:"id"`
	UserID            string    `json:"user_id" db:"user_id"`
	Name              string    `json:"name" db:"name"`
	Price             float64   `json:"price" db:"price"`
	StartingDate      time.Time `json:"starting_date" db:"starting_date"`
	PaymentDate       time.Time `json:"payment_date" db:"payment_date"`
	SubscriptionRenew string    `json:"subscription_renew" db:"subscription_renew"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
