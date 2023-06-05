CREATE SCHEMA IF NOT EXISTS `currency_db`;
USE `currency_db` ;

CREATE TABLE IF NOT EXISTS `currency_db`.`currency` (
  `id`         INT AUTO_INCREMENT NOT NULL,
  `based`      VARCHAR(3) NOT NULL,
  `date`       DATE NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE
);


CREATE TABLE IF NOT EXISTS `currency_db`.`rates` (
  `id`          INT AUTO_INCREMENT NOT NULL,
  `currency`    VARCHAR(3) NOT NULL,
  `rate`        DOUBLE NOT NULL,
  `currency_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE
);