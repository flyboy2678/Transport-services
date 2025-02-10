package store

import (
	"context"
	"database/sql"
)

type Subscription struct {
	ID         int64  `json:"id"`
	User_id    int64  `json:"user_id"`
	Email      string `json:"email"`
	Created_at string `json:"created_at"`
}

type SubscriptionStore struct {
	db *sql.DB
}

func (s *SubscriptionStore) Create(ctx context.Context, subscription *Subscription) error {
	query := `INSERT INTO subscription (user_id, email) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, subscription.User_id, subscription.Email).Scan()
	if err != nil {
		return err
	}

	return nil
}

func (s *SubscriptionStore) DeleteByEmail(ctx context.Context, email string) error {
	query := `DELETE FROM subscription WHERE email = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, email)
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
