package store

import (
	"context"
	"database/sql"
	"errors"
)

var ErrDuplicateEmail = errors.New("a user with that email already exists")

type User struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Phone      string `json:"phone"`
	Created_at string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO "user" (email, password, first_name, last_name, phone)
	VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.Password,
		user.First_name,
		user.Last_name,
		user.Phone,
	).Scan(
		&user.ID,
		&user.Created_at,
	)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique contraint "user_email_key"` {
			return ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	query := `SELECT id, email, password, first_name, last_name, phone, created_at FROM "user" WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}

	err := s.db.QueryRowContext(
		ctx, query, userID,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.First_name,
		&user.Last_name,
		&user.Phone,
		&user.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return user, nil

}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, email, first_name, last_name, password, phone, created_at FROM "user" WHERE email = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}

	err := s.db.QueryRowContext(
		ctx, query, email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.First_name,
		&user.Last_name,
		&user.Phone,
		&user.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *UserStore) UpdateByID(ctx context.Context, user *User) error {
	query := `UPDATE "user" SET email = $1, phone = $2 WHERE id = $3`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx, query, user.Email, user.Phone, user.ID,
	).Scan()

	if err != nil {
		return err
	}

	return nil

}

func (s *UserStore) DeleteByID(ctx context.Context, userID int64) error {
	query := `DELETE FROM "user" WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, userID)
	if err != nil {
		return nil
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
