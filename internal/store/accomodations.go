package store

import (
	"context"
	"database/sql"
)

type Accomodation struct {
	ID              int64   `json:"id"`
	Trip_id         int64   `json:"trip_id"`
	Name            string  `json:"name"`
	Decription      string  `json:"description"`
	Price_per_night float64 `json:"price_per_night"`
	Created_at      string  `json:"created_at"`
}

type AccomodationStore struct {
	db *sql.DB
}

func (s *AccomodationStore) Create(ctx context.Context, accomodation *Accomodation) error {
	query := `
	INSERT INTO accomodation (name, trip_id, description, price_per_night) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		accomodation.Name,
		accomodation.Decription,
		accomodation.Price_per_night,
	).Scan(
		&accomodation.ID,
		&accomodation.Created_at,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccomodationStore) GetByID(ctx context.Context, accomodation_id int64) (*Accomodation, error) {
	query := `SELECET id, trip_id, name, description, price_per_night, created_at 
	FROM accomodation 
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	accomodation := &Accomodation{}

	err := s.db.QueryRowContext(
		ctx, query, accomodation_id,
	).Scan(
		&accomodation.ID,
		&accomodation.Trip_id,
		&accomodation.Name,
		&accomodation.Decription,
		&accomodation.Price_per_night,
		&accomodation.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return accomodation, nil
}

func (s *AccomodationStore) GetByTripID(ctx context.Context, trip_id int64) ([]Accomodation, error) {
	query := `SELECET id, trip_id, name, description, price_per_night, created_at 
	FROM accomodation 
	WHERE trip_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, trip_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accomodations []Accomodation

	for rows.Next() {
		var accomodation Accomodation
		if err := rows.Scan(
			&accomodation.ID,
			&accomodation.Trip_id,
			&accomodation.Name,
			&accomodation.Decription,
			&accomodation.Price_per_night,
			&accomodation.Created_at,
		); err != nil {
			return nil, err
		}
		accomodations = append(accomodations, accomodation)
	}

	return accomodations, nil
}

func (s *AccomodationStore) UpdateByID(ctx context.Context, accomodation *Accomodation) error {
	query := `UPDATE accomodation SET name = $1, description = $2, price_per_night = $3
	WHERE id = $4
	RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx, query, accomodation.Name, accomodation.Decription, accomodation.Price_per_night, accomodation.ID,
	).Scan(&accomodation.ID)

	if err != nil {
		return err
	}

	return nil
}
