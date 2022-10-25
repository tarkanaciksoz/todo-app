DROP DATABASE IF EXISTS `todo-app`;

CREATE DATABASE `todo-app`;

USE `todo-app`;

DROP TABLE IF EXISTS `todos`;

CREATE TABLE `todos` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `value` text,
  `marked` tinyint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;