-- +migrate Up
ALTER TABLE users add column password varchar(255) not null;

-- +migrate Down
ALTEr TABLE users drop column password;