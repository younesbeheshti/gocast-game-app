-- +migrate Up
CREATE TABLE permissions (
                       ID SERIAL PRIMARY KEY,
                       title VARCHAR(191) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
insert into permissions(title) values('user-list');
insert into permissions(title) values('user-delete');


-- +migrate Down
DROP TABLE permissions;