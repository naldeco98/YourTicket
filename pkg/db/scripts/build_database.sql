CREATE TABLE `users`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `role_id` INT NOT NULL,
    `gym_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL
);
CREATE TABLE `gyms`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `address` LONGTEXT NOT NULL,
    `created_at` DATETIME NOT NULL
);
CREATE TABLE `roles`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `role` INT NOT NULL,
    `create_ticket` TINYINT(1) NOT NULL,
    `delete_ticket` TINYINT(1) NOT NULL
);
CREATE TABLE `tickets`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `gym_id` INT NOT NULL,
    `class_id` INT NOT NULL,
    `reservation` DATETIME NOT NULL
);
CREATE TABLE `classes`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `type` VARCHAR(255) NOT NULL,
    `class_date` DATETIME NOT NULL
);
ALTER TABLE
    `users` ADD CONSTRAINT `users_gym_id_foreign` FOREIGN KEY(`gym_id`) REFERENCES `gyms`(`id`);
ALTER TABLE
    `users` ADD CONSTRAINT `users_role_id_foreign` FOREIGN KEY(`role_id`) REFERENCES `roles`(`id`);
ALTER TABLE
    `tickets` ADD CONSTRAINT `tickets_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `users`(`id`);
ALTER TABLE
    `tickets` ADD CONSTRAINT `tickets_gym_id_foreign` FOREIGN KEY(`gym_id`) REFERENCES `gyms`(`id`);
ALTER TABLE
    `tickets` ADD CONSTRAINT `tickets_class_id_foreign` FOREIGN KEY(`class_id`) REFERENCES `classes`(`id`);