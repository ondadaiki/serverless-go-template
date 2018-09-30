package domain

import (
	"database/sql"
	"serverless-go-template/database"
	"time"
)

type User struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ScanUsers(rows *sql.Rows) ([]*User, error) {
	users := []*User{}

	for rows.Next() {
		user, err := ScanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func ScanUser(scanner database.Scanner) (*User, error) {
	var id int64
	var name string
	var createdAt time.Time
	var updatedAt time.Time

	if err := scanner.Scan(
		&id, &name, &createdAt, &updatedAt,
	); err != nil {
		return nil, err
	}

	return &User{
		Id:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
