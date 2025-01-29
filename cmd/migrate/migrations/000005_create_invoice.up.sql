CREATE TABLE IF NOT EXISTS invoice (
    id SERIAL PRIMARY KEY,
    payment_id INT NOT NULL REFERENCES payment(id) ON DELETE CASCADE,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    issue_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    status VARCHAR(21) CHECK (status IN ('paid', 'unpaid', 'overdue')) DEFAULT 'unpaid'
)