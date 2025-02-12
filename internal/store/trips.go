package store

import (
	"context"
	"database/sql"
)

type Trip struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Decription      string  `json:"description"`
	Location        string  `json:"location"`
	Start_date      string  `json:"start_date"`
	End_date        string  `json:"end_date"`
	Price           float64 `json:"price"`
	Seats           int     `json:"seats"`
	Available_seats int     `json:"available_seats"`
	Created_at      string  `json:"created_at"`
}

type TripStore struct {
	db *sql.DB
}

func (s *TripStore) Create(ctx context.Context, trip *Trip) error {
	query := `INSERT INTO trip (name, description, location, start_date, end_date, price, seats, available_seats)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		trip.Name,
		trip.Decription,
		trip.Location,
		trip.Start_date,
		trip.End_date,
		trip.Price,
		trip.Seats,
		trip.Available_seats,
	).Scan(
		&trip.ID,
		&trip.Created_at,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *TripStore) GetByID(ctx context.Context, tripID int64) (*Trip, error) {
	query := `SELECT id, name, description, location, start_date, end_date, price, seats, available_seats, created_at FROM trip WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	trip := &Trip{}

	err := s.db.QueryRowContext(
		ctx, query, tripID,
	).Scan(
		&trip.ID,
		&trip.Name,
		&trip.Decription,
		&trip.Location,
		&trip.Start_date,
		&trip.End_date,
		&trip.Price,
		&trip.Seats,
		&trip.Available_seats,
		&trip.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return trip, nil
}

func (s *TripStore) GetByLocation(ctx context.Context, location string) ([]Trip, error) {
	query := `SELECT id, name, description, location, start_date, end_date, price, seats, available_seats, created_at FROM trip WHERE location = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, location)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []Trip

	for rows.Next() {
		var trip Trip
		if err := rows.Scan(&trip.ID,
			&trip.Name,
			&trip.Decription,
			&trip.Location,
			&trip.Start_date,
			&trip.End_date,
			&trip.Price,
			&trip.Seats,
			&trip.Available_seats,
			&trip.Created_at,
		); err != nil {
			return nil, err
		}

		trips = append(trips, trip)
	}

	return trips, nil
}

func (s *TripStore) GetUpcoming(ctx context.Context) ([]Trip, error) {
	query := `SELECT id, name, description, location, start_date, end_date, price, seats, available_seats, created_at
	FROM trip
	WHERE start_date >= CURRENT_DATE
	ORDER BY start_date ASC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []Trip

	for rows.Next() {
		var trip Trip
		if err := rows.Scan(&trip.ID,
			&trip.Name,
			&trip.Decription,
			&trip.Location,
			&trip.Start_date,
			&trip.End_date,
			&trip.Price,
			&trip.Seats,
			&trip.Available_seats,
			&trip.Created_at,
		); err != nil {
			return nil, err
		}

		trips = append(trips, trip)
	}

	return trips, nil

}

func (s *TripStore) GetAll(ctx context.Context) ([]Trip, error) {
	query := `SELECT id, name, description, location, start_date, end_date, price, seats, available_seats, created_at
	FROM trip
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []Trip

	for rows.Next() {
		var trip Trip
		if err := rows.Scan(&trip.ID,
			&trip.Name,
			&trip.Decription,
			&trip.Location,
			&trip.Start_date,
			&trip.End_date,
			&trip.Price,
			&trip.Seats,
			&trip.Available_seats,
			&trip.Created_at,
		); err != nil {
			return nil, err
		}

		trips = append(trips, trip)
	}

	return trips, nil

}

func (s *TripStore) UpdateByID(ctx context.Context, trip *Trip) error {
	query := `UPDATE trip SET name = $1, description = $2, location = $3, start_date = $4, end_date = $5, price = $6, seats = $7, available_seats = $8
		WHERE id = $9
		RETURING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx, query, trip.Name, trip.Decription, trip.Location, trip.Start_date, trip.End_date, trip.Price, trip.Seats, trip.Available_seats, trip.ID,
	).Scan(&trip.ID)

	if err != nil {
		return err
	}

	return nil
}
