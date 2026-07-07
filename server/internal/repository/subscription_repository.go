package repository

import (
	"context"
	"fmt"
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

func GetAllSubscriptions(pool *pgxpool.Pool, id int) (*[]model.Subscription, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			SELECT * 
			FROM subscriptions
			WHERE user_id = $1
			ORDER BY created_at DESC
	`

	rows, err := pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subscriptions []model.Subscription = []model.Subscription{}

	for rows.Next() {
		var subscription model.Subscription

		err = rows.Scan(
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

		subscriptions = append(subscriptions, subscription)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &subscriptions, nil
}

func GetSubscriptionByID(pool *pgxpool.Pool, id int) (*model.Subscription, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			SELECT id, user_id, name, price, starting_date, payment_date, subscription_renew, created_at, updated_at 
			FROM subscriptions
			WHERE id = $1
	`

	var subscription model.Subscription

	err := pool.QueryRow(ctx, query, id).Scan(
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

	return &subscription, nil
}

func UpdateSubscriptionByID(
	pool *pgxpool.Pool, id int, name string, price float64, isActive bool,
	startingDate time.Time, paymentDate time.Time,
	subscriptionRenew string) (*model.Subscription, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
				UPDATE subscriptions
				SET name = $1, price = $2, is_active = $3, starting_date = $4, payment_date = $5, subscription_renew = $6, updated_at = CURRENT_TIMESTAMP
				WHERE id = $6
				RETURNING id, user_id, name, price, is_active, starting_date, payment_date, subscription_renew, created_at, updated_at
		`

	var subscription model.Subscription

	err := pool.QueryRow(ctx, query, name, price, isActive, startingDate, paymentDate, subscriptionRenew, id).Scan(
		&subscription.ID,
		&subscription.UserID,
		&subscription.Name,
		&subscription.Price,
		&subscription.IsActive,
		&subscription.PaymentDate,
		&subscription.StartingDate,
		&subscription.SubscriptionRenew,
		&subscription.CreatedAt,
		&subscription.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func DeleteSubscription(pool *pgxpool.Pool, id int) error {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			DELETE from subscriptions
			WHERE id = $1
	`

	commandTag, err := pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("User with ID %d not found", id)
	}

	return nil
}
