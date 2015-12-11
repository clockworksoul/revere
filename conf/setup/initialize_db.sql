CREATE DATABASE revere DEFAULT CHARACTER SET = utf8mb4;

use revere;

CREATE TABLE monitors (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  owner VARCHAR(60) NOT NULL DEFAULT '',
  description TEXT NOT NULL DEFAULT '',
  response TEXT NOT NULL DEFAULT '',
  probeType SMALLINT NOT NULL,
  probe TEXT NOT NULL,
  changed DATETIME NOT NULL,
  version INTEGER NOT NULL DEFAULT 0,
  archived DATETIME DEFAULT NULL
) ENGINE = InnoDB;

CREATE TABLE triggers (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  level TINYINT NOT NULL DEFAULT 2,
  triggerOnExit BOOLEAN NOT NULL DEFAULT TRUE,
  periodMs INTEGER NOT NULL DEFAULT 0,
  targetType SMALLINT NOT NULL,
  target TEXT NOT NULL
) ENGINE = InnoDB;

CREATE TABLE monitor_triggers (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  monitor_id INTEGER NOT NULL,
  subprobe TEXT NOT NULL DEFAULT '',
  trigger_id INTEGER NOT NULL,
  FOREIGN KEY (`monitor_id`) REFERENCES monitors(id),
  FOREIGN KEY (`trigger_id`) REFERENCES triggers(id)
) ENGINE = InnoDB;

CREATE TABLE subprobes (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  monitor_id INTEGER NOT NULL,
  name VARCHAR(150) NOT NULL,
  archived DATETIME DEFAULT NULL,
  CONSTRAINT UNIQUE (monitor_id,name)
) ENGINE = InnoDB;

CREATE TABLE subprobe_statuses (
  subprobe_id INTEGER NOT NULL,
  recorded DATETIME NOT NULL,
  state TINYINT NOT NULL,
  silenced BOOLEAN NOT NULL DEFAULT FALSE,
  enteredState DATETIME NOT NULL,
  FOREIGN KEY (`subprobe_id`) REFERENCES subprobes(id),
  INDEX (state, silenced, enteredState, recorded, subprobe_id)
) ENGINE = InnoDB;

CREATE TABLE readings2 (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  subprobe_id INTEGER NOT NULL,
  recorded DATETIME NOT NULL,
  state TINYINT NOT NULL,
  FOREIGN KEY (`subprobe_id`) REFERENCES subprobes(id),
  INDEX (subprobe_id, recorded, id)
) ENGINE = InnoDB;

# Old tables

CREATE TABLE readings (
  id int(11) AUTO_INCREMENT PRIMARY KEY,
  config_id int(11),
  subprobe varchar(255),
  state int(8),
  time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`config_id`) REFERENCES configurations (id) ON DELETE CASCADE
) ENGINE = InnoDB;

CREATE TABLE alerts (
  id int(11) AUTO_INCREMENT PRIMARY KEY,
  config_id int(11),
  subprobe varchar(255),
  time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`config_id`) REFERENCES configurations (id) ON DELETE CASCADE
) ENGINE = InnoDB;

CREATE TABLE silenced_alerts (
  config_id int(11),
  subprobe  varchar(255),
  silence_time timestamp NOT NULL,
  FOREIGN KEY (`config_id`) REFERENCES configurations (id) ON DELETE CASCADE,
  PRIMARY KEY (`config_id`, `subprobe`)
) ENGINE = InnoDB;