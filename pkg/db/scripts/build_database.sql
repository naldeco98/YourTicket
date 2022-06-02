CREATE DATABASE IF NOT EXISTS `testYT`;

USE `testYT`;


DROP TABLE IF EXISTS `roles`;

CREATE TABLE `roles`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `role` VARCHAR(5) NOT NULL,
    `create_ticket` TINYINT(1) NOT NULL,
    `delete_ticket` TINYINT(1) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_UNIQUE` (`id`)
);


DROP TABLE IF EXISTS `gyms`;

CREATE TABLE `gyms`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `address` LONGTEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_UNIQUE` (`id`)
);

DROP TABLE IF EXISTS `classes`;

CREATE TABLE `classes`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `type` VARCHAR(255) NOT NULL,
    `class_date` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_UNIQUE` (`id`)
);

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `role_id` INT NOT NULL,
    `gym_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_UNIQUE` (`id`),
    KEY `fk_users_1_idx` (`role_id`),
    KEY `fk_users_2_idx` (`gym_id`),
    CONSTRAINT `fk_users_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_users_2` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `tickets`;

CREATE TABLE `tickets`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `gym_id` INT NOT NULL,
    `class_id` INT NOT NULL,
    `reservation` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_UNIQUE` (`id`),
    KEY `fk_user_id` (`user_id`),
    KEY `fk_gym_id` (`gym_id`),
    KEY `fk_class_id` (`class_id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_gym_id` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_class_id` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);