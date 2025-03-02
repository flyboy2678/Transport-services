package store

import (
	"context"
	"database/sql"
)

type AccomodationPhoto struct {
	ID              int64  `json:"id"`
	Accomodation_id int64  `json:"accomodation_id"`
	Photo_url       string `json:"photo_url"`
	Created_at      string `json:"created_at"`
}

type AccomodationPhotoStore struct {
	db *sql.DB
}

func (s *AccomodationPhotoStore) Create(ctx context.Context, accomodationPhoto *AccomodationPhoto) error {
	query := `INSERT INTO accomodation_photo (accomodation_id, photo_url)
	RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx, query, accomodationPhoto.Accomodation_id, accomodationPhoto.Photo_url,
	).Scan(&accomodationPhoto.ID, &accomodationPhoto.Created_at)

	if err != nil {
		return err
	}
	return nil
}

func (s *AccomodationPhotoStore) GetByAccomodationId(ctx context.Context, accomodation_id int64) ([]AccomodationPhoto, error) {
	query := `SELECT id, accomodation_id, photo_url, created_at
	FROM accomodation_photo
	WHERE accomodation_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, accomodation_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []AccomodationPhoto
	for rows.Next() {
		var photo AccomodationPhoto
		if err := rows.Scan(photo.ID,
			photo.Photo_url,
			photo.Created_at); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (s *AccomodationPhotoStore) DeleteByID(ctx context.Context, id int64) error {
	query := `DELETE FROM accomodation_photo WHERE id = $1`

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

func (s *AccomodationPhotoStore) DeleteByAccomodationId(ctx context.Context, accomodation_id int64) error {
	query := `DELETE FROM accomodation_photo WHERE accomodation_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, accomodation_id)
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
