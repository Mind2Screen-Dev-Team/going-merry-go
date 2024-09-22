-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `full_name` varchar(150) NOT NULL DEFAULT '' COMMENT 'User Full Name',
    `email` varchar(150) NOT NULL DEFAULT '' COMMENT 'User Email Address',
    `password` text NULL COMMENT 'User Password',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User Create At',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User Updated At',
    `deleted_at` timestamp NULL COMMENT 'User Deleted At',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_email` (`email`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd
