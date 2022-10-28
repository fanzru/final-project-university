CREATE TABLE `papers_users` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` INT(11),
  `paper_name` LONGTEXT,
  `link_pdf` LONGTEXT,
  `is_done` BOOLEAN,
  `created_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_users_papers` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE NO ACTION
);