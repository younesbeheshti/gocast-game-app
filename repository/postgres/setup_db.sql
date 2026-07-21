CREATE TABLE users (
                       ID SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       phone_number VARCHAR(255) NOT NULL UNIQUE,
                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);