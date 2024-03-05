CREATE TABLE IF NOT EXISTS cars (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name Varchar(50) NOT NULL,
    brand Varchar(20) NOT NULL,
    model Varchar(30) NOT NULL,
    hourse_power INTEGER DEFAULT 0,
    colour VARCHAR(20) NOT NULL DEFAULT 'black',
    engine_cap DECIMAL(10,2) NOT NULL DEFAULT 1.0,
    year INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS customers(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    gmail VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

ALTER TABLE customers
ADD CONSTRAINT must_have UNIQUE (deleted_at,phone);