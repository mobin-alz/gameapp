-- +migrate Up
-- MYSQL 8.0 set the role value to `user` for all old records (order priority)
-- TODO find a better solution instead of keeping the order !!!
ALTER TABLE `users` ADD COLUMN `role` ENUM('user', 'admin') NOT NULL ;

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;
