CREATE TABLE Category (
    Id SERIAL PRIMARY KEY,
    Name TEXT NOT NULL,
    Created_at TIMESTAMP DEFAULT NOW(),
    Updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE Product (
    Id SERIAL PRIMARY KEY,
    Name TEXT NOT NULL,
    Price DECIMAL NOT NULL,
    Category_id INT,
    Created_at TIMESTAMP DEFAULT NOW(),
    Updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (Category_id) REFERENCES Category(Id)
);