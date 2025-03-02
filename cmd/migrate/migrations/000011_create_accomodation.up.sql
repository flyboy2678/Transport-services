CREATE TABLE IF NOT EXISTS accomodation (
    id SERIAL PRIMARY KEY,
    trip_id INT NOT NULL REFERENCES trip(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price_per_night FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
)
