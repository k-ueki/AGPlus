-- +goose Up
CREATE TABLE IF NOT EXISTS user (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  firebase_uid VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  -- photo_url VARCHAR(255),
  department_id int(10) UNSIGNED NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY `firebase_uid` (`firebase_uid`),
  FOREIGN KEY (department_id) REFERENCES faculty(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE user;
