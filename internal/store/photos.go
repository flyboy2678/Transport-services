package store

import (
	"context"
	"database/sql"
)

type Photo struct {
	ID          int64  `json:"id"`
	Trip_id     int64  `json:"trip_id"`
	Photo_url   string `json:"photo_url"`
	Uploaded_at string `json:"uploaded_at"`
}

type PhotoStore struct {
	db *sql.DB
}

func (s *PhotoStore) Create(ctx context.Context, photo *Photo) error {
	query := `INSERT INTO photo (trip_id, photo_url) 
	VALUES ($1, $2)
	RETURNING id, uploaded_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		photo.Trip_id,
		photo.Photo_url).Scan(&photo.ID, &photo.Uploaded_at)
	if err != nil {
		return err
	}

	return nil
}

func (s *PhotoStore) GetByID(ctx context.Context, photoID int64) (*Photo, error) {
	query := `SELECT id, trip_id, photo_url, uploaded_at FROM photo WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	photo := &Photo{}

	err := s.db.QueryRowContext(ctx, query, photoID).Scan(
		&photo.ID, &photo.Trip_id, &photo.Photo_url, &photo.Uploaded_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return photo, nil
}

func (s *PhotoStore) GetByTripID(ctx context.Context, tripID int64) ([]Photo, error) {
	query := `SELECT id, trip_id, photo_url, uploaded_at FROM photo WHERE trip_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.ID, &photo.Trip_id, &photo.Photo_url, &photo.Uploaded_at); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return photos, nil

}

func (s *PhotoStore) DeleteByID(ctx context.Context, id int64) error {
	query := `DELETE FROM photo WHERE id = $1`

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

func (s *PhotoStore) DeleteByTripID(ctx context.Context, tripID int64) error {
	query := `DELETE FROM photo WHERE trip_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, tripID)
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
