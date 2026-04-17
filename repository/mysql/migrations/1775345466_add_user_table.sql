-- +migrate Up
CREATE TABLE `users` (
                         id int PRIMARY KEY AUTO_INCREMENT,
                         name varchar(191),
                         phone_number varchar(191) not null UNIQUE,
                         create_at timestamp DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE `users`;