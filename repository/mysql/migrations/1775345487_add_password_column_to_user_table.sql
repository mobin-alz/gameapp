-- +migrate Up
ALTER TABLE users add column password varchar(191) not null;

-- +migrate Down
ALTEr TABLE users drop column password;