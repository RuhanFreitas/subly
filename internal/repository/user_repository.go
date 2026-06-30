package repository

import (
	"context"
	"subly/internal/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(pool *pgxpool.Pool, user *model.User) (*model.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			INSERT INTO users (first_name, last_name, email, password)
			VALUES ($1, $2, $3, $4)
			RETURNING id, first_name, last_name, email, created_at, updated_at
	`

	var err error = pool.QueryRow(ctx, query, user.FirstName, user.LastName, user.Email, user.Password).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
