package store

import (
	"context"
	"database/sql"
)

type Invoice struct {
	ID             int64  `json:"id"`
	Payment_id     int64  `json:"payment_id"`
	Invoice_number string `json:"invoice_number"`
	Issue_at       string `json:"issue_at"`
	Due_date       string `json:"due_date"`
	Status         string `json:"status"`
}

type InvoiceStore struct {
	db *sql.DB
}

func (s *InvoiceStore) Create(ctx context.Context, invoice *Invoice) error {
	query := `INSERT INTO invoice (payment_id, invoice_number, issue_at, due_date, status)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, invoice.Payment_id, invoice.Invoice_number, invoice.Issue_at, invoice.Due_date, invoice.Status).Scan(&invoice.ID, &invoice.Issue_at)
	if err != nil {
		return err
	}
	return nil
}

func (s *InvoiceStore) UpdateByInvoiceNumber(ctx context.Context, invoice *Invoice) error {
	query := `UPDATE invoice 
	SET status = $1 
	WHERE invoice_number = $2
	RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, invoice.Status, invoice.Invoice_number).Scan(&invoice.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *InvoiceStore) GetByInvoiceNumber(ctx context.Context, invoiceNumber string) (*Invoice, error) {
	query := `SELECT id, payment_id, invoice_number, issue_at, due_date, status FROM invoice WHERE invoice_number = $1 `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	invoice := &Invoice{}

	err := s.db.QueryRowContext(ctx, query, invoiceNumber).Scan(
		&invoice.ID, &invoice.Payment_id, &invoice.Issue_at, &invoice.Due_date, &invoice.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return invoice, nil
}
