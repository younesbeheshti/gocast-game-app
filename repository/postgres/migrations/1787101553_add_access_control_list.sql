-- +migrate Up

CREATE TYPE actor AS ENUM('role', 'user');

CREATE TABLE access_controls (
                       ID SERIAL PRIMARY KEY,
                       actor_id VARCHAR(191) NOT NULL UNIQUE,
                       actor_type actor NOT NULL,
                       permission_id INT NOT NULL REFERENCES permissions(ID),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

-- +migrate Down
DROP TABLE access_controls;