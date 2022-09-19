CREATE TABLE center_skills (
  id int NOT NULL,
  name varchar(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE color_types (
  id int NOT NULL,
  name varchar(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE groups (
  id int NOT NULL ,
  name varchar(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE lives (
  id int NOT NULL ,
  name varchar(100) NOT NULL,
  group_id int NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT Live_ibfk_1 FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE members (
  id int NOT NULL ,
  name varchar(100) NOT NULL,
  first_name varchar(100) DEFAULT NULL,
  group_id int NOT NULL,
  phase int NOT NULL,
  graduated int NOT NULL DEFAULT '0',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT Member_ibfk_1 FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE music (
  id int NOT NULL ,
  name varchar(100) NOT NULL,
  normal int NOT NULL,
  pro int NOT NULL,
  master int NOT NULL,
  length int NOT NULL,
  color_type_id int NOT NULL,
  live_id int NOT NULL,
  pro_plus int DEFAULT NULL,
  music_bonus int DEFAULT '0',
  setlist_id int DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT Music_ibfk_1 FOREIGN KEY (live_id) REFERENCES lives (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT Music_ibfk_2 FOREIGN KEY (color_type_id) REFERENCES color_types (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE photograph (
  id int NOT NULL ,
  name varchar(100) NOT NULL,
  group_id int NOT NULL,
  abbreviation varchar(10) NOT NULL DEFAULT '',
  photo_type varchar(10) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT Photograph_ibfk_1 FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE scenes (
  id int NOT NULL ,
  photograph_id int NOT NULL,
  member_id int NOT NULL,
  color_type_id int NOT NULL,
  vocal_max int NOT NULL,
  dance_max int NOT NULL,
  peformance_max int NOT NULL,
  center_skill_name varchar(100)  DEFAULT NULL,
  expected_value varchar(5)  DEFAULT NULL,
  ssr_plus int NOT NULL DEFAULT '0',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT Scene_ibfk_2 FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT Scene_ibfk_3 FOREIGN KEY (photograph_id) REFERENCES photograph (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT Scene_ibfk_5 FOREIGN KEY (color_type_id) REFERENCES color_types (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE skills (
  id int NOT NULL ,
  name varchar(100) NOT NULL,
  combo_up_percent int DEFAULT NULL,
  duration_sec int DEFAULT NULL,
  expected_value double NOT NULL,
  interval_sec int NOT NULL,
  occurrence_percent int NOT NULL,
  score_up_percent int DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE producers (
  id int NOT NULL ,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE producer_scenes (
  id int NOT NULL ,
  producer_id INT NOT NULL,
  photograph_id INT NOT NULL,
  member_id INT NOT NULL,
  have INT DEFAULT 0,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_producer_scenes_producer_id FOREIGN KEY (producer_id) REFERENCES producers (id),
  CONSTRAINT fk_producer_scenes_photo_id FOREIGN KEY (photograph_id) REFERENCES photograph (id),
  CONSTRAINT fk_producer_scenes_member_id FOREIGN KEY (member_id) REFERENCES members (id)
);

CREATE TABLE producer_members (
  id int NOT NULL ,
  producer_id INT NOT NULL,
  member_id INT NOT NULL,
  bond_level_curent INT NOT NULL,
  bond_level_collection_max INT NOT NULL,
  bond_level_scene_max INT NOT NULL,
  discography_disc_total INT NOT NULL,
  discography_disc_total_max INT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_producer_members_producer_id FOREIGN KEY (producer_id) REFERENCES producers (id),
  CONSTRAINT fk_producer_members_member_id FOREIGN KEY (member_id) REFERENCES members (id)
);

CREATE TABLE producer_offices (
  id int NOT NULL ,
  producer_id INT NOT NULL,
  office_bonus INT DEFAULT 0,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_producer_offices_producer_id FOREIGN KEY (producer_id) REFERENCES producers (id)
);
