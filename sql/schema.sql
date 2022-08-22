CREATE TABLE center_skills (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE color_types (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE lives (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `group_id` int NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `group_id` (`group_id`),
  CONSTRAINT `Live_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE members (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  first_name varchar(100) DEFAULT NULL,
  `group_id` int NOT NULL,
  `phase` int NOT NULL,
  `graduated` tinyint(1) NOT NULL DEFAULT '0',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `group_id` (`group_id`),
  CONSTRAINT `Member_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE music (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `normal` int NOT NULL,
  `pro` int NOT NULL,
  `master` int NOT NULL,
  `length` int NOT NULL,
  `color_type_id` int NOT NULL,
  `live_id` int NOT NULL,
  `pro_plus` int NOT NULL,
  `music_bonus` tinyint(1) NOT NULL,
  `setlist_id` int NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `live_id` (`live_id`),
  KEY `color_type_id` (`color_type_id`),
  CONSTRAINT `Music_ibfk_1` FOREIGN KEY (`live_id`) REFERENCES `lives` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Music_ibfk_2` FOREIGN KEY (`color_type_id`) REFERENCES `color_types` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE photograph (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `group_id` int NOT NULL,
  `abbreviation` varchar(191) NOT NULL DEFAULT '',
  `type` varchar(191) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `group_id` (`group_id`),
  CONSTRAINT `Photograph_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE scenes (
  `id` int NOT NULL AUTO_INCREMENT,
  `photograph_id` int NOT NULL,
  `member_id` int NOT NULL,
  `color_type_id` int NOT NULL,
  `vocal_max` int NOT NULL,
  `dance_max` int NOT NULL,
  `peformance_max` int NOT NULL,
  `center_skill_name` varchar(100)  DEFAULT NULL,
  `expected_value` varchar(5)  DEFAULT NULL,
  `ssr_plus` tinyint(1) NOT NULL DEFAULT '0',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `member_id` (`member_id`),
  KEY `photograph_id` (`photograph_id`),
  KEY `color_type_id` (`color_type_id`),
  CONSTRAINT `Scene_ibfk_2` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Scene_ibfk_3` FOREIGN KEY (`photograph_id`) REFERENCES `photograph` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Scene_ibfk_5` FOREIGN KEY (`color_type_id`) REFERENCES `color_types` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE skills (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `combo_up_percent` int DEFAULT NULL,
  `duration_sec` int DEFAULT NULL,
  `expected_value` double NOT NULL,
  `interval_sec` int NOT NULL,
  `occurrence_percent` int NOT NULL,
  `score_up_percent` int DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE producers (
  `id` int NOT NULL AUTO_INCREMENT,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE producer_scenes (
  `id` int NOT NULL AUTO_INCREMENT,
  producer_id INT NOT NULL,
  photograph_id INT NOT NULL,
  member_id INT NOT NULL,
  have TINYINT(1) DEFAULT 0,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_producer_scenes_producer_id` FOREIGN KEY (`producer_id`) REFERENCES `producers` (`id`),
  CONSTRAINT `fk_producer_scenes_photo_id` FOREIGN KEY (`photograph_id`) REFERENCES `photograph` (`id`),
  CONSTRAINT `fk_producer_scenes_member_id` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;