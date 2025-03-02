package store

import (
	"context"
	"database/sql"
)

type ActivityPhoto struct {
	ID          int64  `json:"id"`
	Activity_id int64  `json:"activity_id"`
	Photo_url   string `json:"photo_url"`
	Uploaded_at string `json:"uploaded_at"`
}

type ActivityPhotoStore struct {
	db *sql.DB
}

func (s *ActivityPhotoStore) Create(ctx context.Context, photo *ActivityPhoto) error {
	query := `INSERT INTO activity_photo (activity_id, photo_url)
	VALUES ($1, $2) 
	RETURNING id, uploaded_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, photo.ID, photo.Photo_url).Scan(&photo.ID, &photo.Uploaded_at)
	if err != nil {
		return err
	}

	return nil
}

func (s *ActivityPhotoStore) GetById(ctx context.Context, id int64) (*ActivityPhoto, error) {
	query := `SELECT id, activity_id, photo_url, uploaded_at 
	FROM activity_photo 
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	photo := &ActivityPhoto{}
	err := s.db.QueryRowContext(
		ctx, query, id,
	).Scan(
		&photo.ID,
		&photo.Activity_id,
		&photo.Photo_url,
		&photo.Uploaded_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return photo, nil
}

func (s *ActivityPhotoStore) GetByActivityId(ctx context.Context, activity_id int64) ([]ActivityPhoto, error) {
	query := `SELECT id, activity_id, photo_url, uploaded_at 
	FROM activity_photo 
	WHERE activity_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, activity_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []ActivityPhoto

	for rows.Next() {
		var photo ActivityPhoto
		if err := rows.Scan(
			photo.ID,
			photo.Activity_id,
			photo.Photo_url,
			photo.Uploaded_at,
		); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (s *ActivityPhotoStore) DeleteById(ctx context.Context, id int64) error {
	query := `DELETE FROM activity_photo WHERE id = $1`

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

func (s *ActivityPhotoStore) DeleteByActivityId(ctx context.Context, id int64) error {
	query := `DELETE FROM activity_photo WHERE activity_id = $1`

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
