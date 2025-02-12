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
	RETURING id, created_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx, query, payment.Booking_id, payment.User_id, payment.Amount, payment.Status, payment.Transaction_id,
	).Scan(
		&payment.ID, &payment.Created_at,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentStore) GetByUserID(ctx context.Context, userId int64) ([]Payment, error) {
	query := `
	SELECT id, booking_id, user_id, amount, status, transaction_id, created_at
	FROM payment
	WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment

	for rows.Next() {
		var payment Payment
		if err := rows.Scan(
			payment.ID,
			payment.Booking_id,
			payment.User_id,
			payment.Amount,
			payment.Status,
			payment.Transaction_id,
			payment.Created_at,
		); err != nil {
			return nil, err
		}

		payments = append(payments, payment)

	}

	return payments, nil

}

func (s *PaymentStore) UpdateByID(ctx context.Context, payment *Payment) error {
	query := `
	UPDATE payment SET  status = $1
	WHERE id = $2
	RETURNING id
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, payment.Status, payment.ID).Scan()
	if err != nil {
		return err
	}
	return nil
}
