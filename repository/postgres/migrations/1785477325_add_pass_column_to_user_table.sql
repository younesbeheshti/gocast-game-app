-- +migrate Up
ALTER TABLE users add column password VARCHAR(255) NOT NULL;


-- +migrate Down
ALTER TABLE users drop column password;
