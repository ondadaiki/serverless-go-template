package repository

import (
	"context"
	"database/sql"
	"serverless-go-template/domain"
)

func FindUser(db *sql.DB, ctx context.Context, id string) (*domain.User, error) {
	user, err := domain.ScanUser(db.QueryRowContext(ctx,
		`SELECT * FROM users WHERE id = $1`,
		id,
	))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers(db *sql.DB, ctx context.Context) ([]*domain.User, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT * FROM users`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users, err := domain.ScanUsers(rows)
	if err != nil {
		return nil, err
	}

	return users, nil
}
