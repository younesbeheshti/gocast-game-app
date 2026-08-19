-- +migrate Up
CREATE TABLE permissions (
                       ID SERIAL PRIMARY KEY,
                       title VARCHAR(191) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE permissions;