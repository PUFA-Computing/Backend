CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY NOT NULL,
    image VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)