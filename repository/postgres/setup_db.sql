CREATE TABLE users (
                       ID SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       phone_number VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);