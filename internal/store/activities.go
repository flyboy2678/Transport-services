package store

import (
	"context"
	"database/sql"
)

type Activity struct {
	ID          int64   `json:"id"`
	Trip_id     int64   `json:"trip_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Created_at  string  `json:"created_at"`
}

type ActivityStore struct {
	db *sql.DB
}

func (s *ActivityStore) Create(ctx context.Context, activity *Activity) error {
	query := `INSERT INTO activity (trip_id, name, description, price)
	VALUES ($1, $2, $3)
	RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, activity.Trip_id, activity.Name, activity.Description, activity.Price).Scan(
		&activity.ID,
		&activity.Created_at,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *ActivityStore) GetById(ctx context.Context, id int64) (*Activity, error) {
	query := `SELECT id, trip_id, name, description, price, created_at
	FROM activity 
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	activty := &Activity{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&activty.ID,
		&activty.Trip_id,
		&activty.Name,
		&activty.Description,
		&activty.Price,
		&activty.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return activty, nil
}

func (s *ActivityStore) GetByTripId(ctx context.Context, tripId int64) ([]Activity, error) {
	query := `SELECT id, trip_id, name, description, price, created_at
	FROM activity
	WHERE trip_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, tripId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var activities []Activity

	for rows.Next() {
		var activity Activity
		if err := rows.Scan(
			activity.ID,
			activity.Trip_id,
			activity.Name,
			activity.Description,
			activity.Price,
			activity.Created_at,
		); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

func (s *ActivityStore) UpdateById(ctx context.Context, activity *Activity) error {
	query := `UPDATE activity
	SET name = $1, description = $2, price = $3
	WHERE id = $4
	RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, activity.Name, activity.Description, activity.Price, activity.ID).Scan(&activity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ActivityStore) DeleteById(ctx context.Context, id int64) error {
	query := `DELETE FROM activity WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, id)
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
