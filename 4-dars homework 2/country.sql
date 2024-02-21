CREATE TABLE IF NOT EXISTS countries(
    id uuid PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    code INTEGER UNIQUE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);