-- +migrate Up
CREATE TABLE users (
                       ID SERIAL PRIMARY KEY,
                       name VARCHAR(191) NOT NULL,
                       phone_number VARCHAR(191) NOT NULL UNIQUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE users;