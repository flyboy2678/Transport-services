package store

import (
	"context"
	"database/sql"
)

type Booking struct {
	ID         int64  `json:"id"`
	User_id    int64  `json:"user_id"`
	Trip_id    int64  `json:"trip_id"`
	Status     string `json:"status"`
	Created_at string `json:"created_at"`
}

type BookingStore struct {
	db *sql.DB
}

func (s *BookingStore) Create(ctx context.Context, booking *Booking) error {
	query := `INSERT INTO booking (user_id, trip_id, status)
	VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx, query, booking.User_id, booking.Trip_id, booking.Status,
	).Scan()
	if err != nil {
		return err
	}

	return nil
}

func (s *BookingStore) GetByID(ctx context.Context, bookingID int64) (*Booking, error) {
	query := `SELECT id, user_id, trip_id, status, created_at FROM booking WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	booking := &Booking{}

	err := s.db.QueryRowContext(ctx, query, bookingID).Scan(
		&booking.ID,
		&booking.User_id,
		&booking.Trip_id,
		&booking.Status,
		&booking.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return booking, nil
}

func (s *BookingStore) GetByTripID(ctx context.Context, tripID int64) ([]Booking, error) {
	query := `SELECT id, user_id, trip_id, status, created_at FROM booking WHERE trip_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, tripID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking

	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.User_id, &booking.Trip_id, &booking.Status, &booking.Created_at); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingStore) GetByUserID(ctx context.Context, userID int64) ([]Booking, error) {
	query := `SELECT id, user_id, trip_id, status, created_at FROM booking WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking

	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.ID, &booking.User_id, &booking.Trip_id, &booking.Status, &booking.Created_at); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingStore) UpdateByID(ctx context.Context, booking *Booking) error {
	query := `UPDATE booking SET user_id = $1, trip_id = $2, status = $3 WHERE id = $4`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, booking.User_id, booking.Trip_id, booking.Status, booking.ID).Scan()
	if err != nil {
		return err
	}

	return nil
}
