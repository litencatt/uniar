CREATE TABLE center_skills (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE color_types (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE groups (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE lives (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE members (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  first_name varchar(100) DEFAULT NULL,
  group_id integer NOT NULL,
  phase integer NOT NULL,
  graduated integer NOT NULL DEFAULT '0',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT Member_ibfk_1 FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE music (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  normal integer NOT NULL,
  pro integer NOT NULL,
  master integer NOT NULL,
  length integer NOT NULL,
  color_type_id integer NOT NULL,
  live_id integer NOT NULL,
  pro_plus integer DEFAULT NULL,
  music_bonus integer DEFAULT '0',
  setlist_id integer DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT Music_ibfk_1 FOREIGN KEY (live_id) REFERENCES lives (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT Music_ibfk_2 FOREIGN KEY (color_type_id) REFERENCES color_types (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE photograph (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  group_id integer NOT NULL,
  abbreviation varchar(10) NOT NULL DEFAULT '',
  photo_type varchar(10) NOT NULL,
  released_at DATETIME NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT Photograph_ibfk_1 FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE scenes (
  id integer PRIMARY KEY AUTOINCREMENT,
  photograph_id integer NOT NULL,
  member_id integer NOT NULL,
  color_type_id integer NOT NULL,
  vocal_max integer NOT NULL,
  dance_max integer NOT NULL,
  performance_max integer NOT NULL,
  center_skill_name varchar(100)  DEFAULT NULL,
  expected_value varchar(100)  DEFAULT NULL,
  ssr_plus integer NOT NULL DEFAULT '0',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT Scene_ibfk_2 FOREIGN KEY (member_id) REFERENCES members (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT Scene_ibfk_3 FOREIGN KEY (photograph_id) REFERENCES photograph (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT Scene_ibfk_5 FOREIGN KEY (color_type_id) REFERENCES color_types (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE skills (
  id integer PRIMARY KEY AUTOINCREMENT,
  name varchar(100) NOT NULL,
  combo_up_percent integer DEFAULT NULL,
  duration_sec integer DEFAULT NULL,
  expected_value double NOT NULL,
  interval_sec integer NOT NULL,
  occurrence_percent integer NOT NULL,
  score_up_percent integer DEFAULT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE producers (
  id integer PRIMARY KEY AUTOINCREMENT,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE producer_scenes (
  producer_id integer NOT NULL,
  scene_id integer NOT NULL,
  have integer NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(producer_id, scene_id),
  CONSTRAINT fk_producer_scenes_producer_id FOREIGN KEY (producer_id) REFERENCES producers (id),
  CONSTRAINT fk_producer_scenes_scene_id FOREIGN KEY (scene_id) REFERENCES scenes (id)
);

CREATE TABLE producer_members (
  id integer PRIMARY KEY AUTOINCREMENT,
  producer_id integer NOT NULL,
  member_id integer NOT NULL,
  bond_level_curent integer NOT NULL,
  bond_level_collection_max integer NOT NULL,
  bond_level_scene_max integer NOT NULL,
  discography_disc_total integer NOT NULL,
  discography_disc_total_max integer NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_producer_members_producer_id FOREIGN KEY (producer_id) REFERENCES producers (id),
  CONSTRAINT fk_producer_members_member_id FOREIGN KEY (member_id) REFERENCES members (id)
);

CREATE TABLE producer_offices (
  id integer PRIMARY KEY AUTOINCREMENT,
  producer_id integer NOT NULL,
  office_bonus integer DEFAULT 0,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_producer_offices_producer_id FOREIGN KEY (producer_id) REFERENCES producers (id)
);
