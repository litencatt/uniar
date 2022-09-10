CREATE TABLE `CenterSkill` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `ColorType` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Live` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `group_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `group_id` (`group_id`),
  CONSTRAINT `Live_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `Group` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Member` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `group_id` int NOT NULL,
  `graduated` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `group_id` (`group_id`),
  CONSTRAINT `Member_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `Group` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Music` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `normal` int NOT NULL,
  `pro` int NOT NULL,
  `master` int NOT NULL,
  `length` int NOT NULL,
  `color_type_id` int NOT NULL,
  `live_id` int NOT NULL,
  `pro_plus` int NOT NULL,
  `music_bonus` tinyint(1) NOT NULL,
  `setlist_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `live_id` (`live_id`),
  KEY `color_type_id` (`color_type_id`),
  CONSTRAINT `Music_ibfk_1` FOREIGN KEY (`live_id`) REFERENCES `Live` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Music_ibfk_2` FOREIGN KEY (`color_type_id`) REFERENCES `ColorType` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Photograph` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `group_id` int NOT NULL,
  `abbreviation` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `type` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `group_id` (`group_id`),
  CONSTRAINT `Photograph_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `Group` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Scene` (
  `id` int NOT NULL AUTO_INCREMENT,
  `color_type_id` int NOT NULL,
  `dance_max` int NOT NULL,
  `member_id` int NOT NULL,
  `peformance_max` int NOT NULL,
  `photograph_id` int NOT NULL,
  `vocal_max` int NOT NULL,
  `center_skill_name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `skill_name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ssr_plus` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `member_id` (`member_id`),
  KEY `photograph_id` (`photograph_id`),
  KEY `color_type_id` (`color_type_id`),
  CONSTRAINT `Scene_ibfk_2` FOREIGN KEY (`member_id`) REFERENCES `Member` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Scene_ibfk_3` FOREIGN KEY (`photograph_id`) REFERENCES `Photograph` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Scene_ibfk_5` FOREIGN KEY (`color_type_id`) REFERENCES `ColorType` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Skill` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `combo_up_percent` int DEFAULT NULL,
  `duration_sec` int DEFAULT NULL,
  `expected_value` double NOT NULL,
  `interval_sec` int NOT NULL,
  `occurrence_percent` int NOT NULL,
  `score_up_percent` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
