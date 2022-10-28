CREATE TABLE `sentences_labels` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `paper_id` INT(11),
  `head` LONGTEXT,
  `text` LONGTEXT,
  `is_important` BOOLEAN,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_paper_sentences` FOREIGN KEY (`paper_id`) REFERENCES `papers_users`(`id`) ON DELETE CASCADE ON UPDATE NO ACTION
);