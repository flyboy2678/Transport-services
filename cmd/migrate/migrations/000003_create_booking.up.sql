CREATE TABLE IF NOT EXISTS booking (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES "user"(id) on DELETE CASCADE,
    trip_id INT NOT  NULL REFERENCES trip(id) on DELETE CASCADE,
    status VARCHAR(20) CHECK (status IN ('pending', 'confirmed','cancelled')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, trip_id)
)