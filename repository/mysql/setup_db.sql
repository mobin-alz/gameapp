CREATE TABLE `users` (
                         id int PRIMARY KEY AUTO_INCREMENT,
                         name varchar(255),
                         phone_number varchar(255) not null UNIQUE,
                         password varchar(255) not null,
                         create_at timestamp DEFAULT CURRENT_TIMESTAMP
)