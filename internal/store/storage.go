package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}
	//add more interface like based on the tables we are
	// on having in our database
}

func NewStorage(db *sql.DB) Storage{
	return Storage{
		Users: &UsersStore{db},
	}
}