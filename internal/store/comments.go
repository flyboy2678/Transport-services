package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID         int64  `json:"id"`
	User_id    int64  `json:"user_id"`
	Trip_id    int64  `json:"trip_id"`
	Comment    string `json:"comment"`
	Rating     int    `json:"rating"`
	Created_at string `json:"created_at"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `INSERT INTO comment (user_id, trip_id, comment, rating)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, comment.User_id, comment.Trip_id, comment.Comment, comment.Rating).Scan(
		&comment.ID, &comment.Created_at,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) GetByID(ctx context.Context, commentID int64) (*Comment, error) {
	query := `SELECT id, user_id, trip_id, comment, rating FROM comment WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	comment := &Comment{}

	err := s.db.QueryRowContext(ctx, query, commentID).Scan(
		&comment.ID, &comment.User_id, &comment.Trip_id, &comment.Comment, &comment.Rating, &comment.Created_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return comment, nil
}

func (s *CommentStore) GetByTripID(ctx context.Context, tripID int64) ([]Comment, error) {
	query := `SELECT id, user_id, trip_id, comment, rating FROM comment WHERE trip_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(
			&comment.ID, &comment.User_id, &comment.Trip_id, &comment.Comment, &comment.Rating, &comment.Created_at,
		); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (s *CommentStore) DeleteByID(ctx context.Context, commentID int64) error {
	query := `DELETE FROM comment WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, commentID)
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

func (s *CommentStore) DeleteByTripID(ctx context.Context, tripID int64) error {
	query := `DELETE FROM comment WHERE trip_id = $1`

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
