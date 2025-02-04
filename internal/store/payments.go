package store

import (
	"context"
	"database/sql"
)

type Payment struct {
	ID             int64   `json:"id"`
	Booking_id     int64   `json:"booking_id"`
	User_id        int64   `json:"user_id"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
	Transaction_id string  `json:"transation_id"`
	Created_at     string  `json:"created_at"`
}

type PaymentStore struct {
	db *sql.DB
}

func (s *PaymentStore) Create(ctx context.Context, payment *Payment) error {
	query := `INSERT INTO payment (booking_id, user_id, amount, status, transaction_id)
	VALUES ($1, $2, $3, $4, $5)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, payment.Booking_id, payment.User_id, payment.Amount, payment.Status, payment.Transaction_id).Scan()
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentStore) UpdateByID(ctx context.Context, payment *Payment) error {
	query := `
	UPDATE payment SET booking_id = $1, user_id = $2, amount = $3, status = $4, transaction_id = $5 
	WHERE id = $6
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, payment.Booking_id, payment.User_id, payment.Amount, payment.Status, payment.Transaction_id, payment.ID).Scan()
	if err != nil {
		return err
	}
	return nil
}
