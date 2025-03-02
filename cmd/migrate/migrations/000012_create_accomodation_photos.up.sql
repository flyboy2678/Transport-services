CREATE TABLE IF NOT EXISTS accomodation_photo (
    id SERIAL PRIMARY KEY,
    accomodation_id INT NOT NULL REFERENCES accomodation(id) ON DELETE CASCADE,
    photo_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
