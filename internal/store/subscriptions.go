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
	query := `INSERT INTO subscription (user_id, email) VALUES ($1, $2)
	RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, subscription.User_id, subscription.Email).Scan(
		&subscription.ID, &subscription.Created_at,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *SubscriptionStore) GetAll(ctx context.Context) ([]Subscription, error) {
	query := `SELECT id, user_id, email, created_at FROM subscription`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []Subscription
	for rows.Next() {
		var sub Subscription
		if err := rows.Scan(&sub.ID, &sub.User_id, &sub.Email, &sub.Created_at); err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}

	return subs, nil
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

func (s *SubscriptionStore) DeleteByUserID(ctx context.Context, userID int64) error {
	query := `DELETE FROM subscription WHERE user_id = $1`

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
