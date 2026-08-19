-- +migrate Up
CREATE TYPE role_type AS ENUM('user', 'admin');
ALTER TABLE users ADD COLUMN role role_type NOT NULL DEFAULT 'user';


-- +migrate Down
ALTER TABLE users DROP COLUMN role;
