package repository

import (
	"context"
	"subly/internal/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateSubscription(pool *pgxpool.Pool, subscription *model.Subscription) (*model.Subscription, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			INSERT INTO subscriptions (user_id, name, price, starting_date, payment_date, subscription_renew)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, user_id, name, price, starting_date, payment_date, subscription_renew, created_at, updated_at
	`

	err := pool.QueryRow(
		ctx, query, subscription.UserID, subscription.Name, subscription.Price,
		subscription.StartingDate, subscription.PaymentDate, subscription.SubscriptionRenew,
	).Scan(
		&subscription.ID,
		&subscription.UserID,
		&subscription.Name,
		&subscription.Price,
		&subscription.StartingDate,
		&subscription.PaymentDate,
		&subscription.SubscriptionRenew,
		&subscription.CreatedAt,
		&subscription.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
