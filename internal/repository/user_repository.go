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

func GetUserByEmail(pool *pgxpool.Pool, email string) (*model.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			SELECT id, first_name, last_name, email, password, created_at, updated_at
			FROM users
			WHERE email = $1
	`

	var user model.User

	err := pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByID(pool *pgxpool.Pool, id int) (*model.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			SELECT id, first_name, last_name, email, password, created_at, updated_at
			FROM users
			WHERE id = $1
	`

	var user model.User

	err := pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUserByID(pool *pgxpool.Pool, id int, firstName string, lastName string, email string, password string) (*model.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
			UPDATE users
			SET first_name = $1, last_name = $2, email = $3, password = $4, updated_at = CURRENT_TIMESTAMP
			WHERE id = $4
			RETURNING id, first_name, last_name, email, created_at, updated_at
	`

	var user model.User

	err := pool.QueryRow(ctx, query, firstName, lastName, email, password, id).Scan(
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

	return &user, nil
}
