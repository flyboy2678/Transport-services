package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	First_name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Phone      string `json:"phone"`
	Created_at string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s* UsersStore) Create(ctx context.Context, user *User) error{
	query := `
	INSERT INTO users (email, password, first_name, last_name, phone)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at
	`
	err := s.db.QueryRowContext(
		ctx, 
		query,
		user.Email,
		user.Password,
		user.First_name,
		user.Last_Name,
		user.Phone,
	).Scan(
		&user.ID,
		&user.Created_at,
	)
	if err!= nil {
		return  err		
	}
	
	return nil
}