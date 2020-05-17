-- +goose Up
CREATE TABLE IF NOT EXISTS class (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  semester VARCHAR(55) NOT NULL,
  credits INT(10) UNSIGNED NOT NULL,
  teacher VARCHAR(255) NOT NULL,
  description LONG,
  course_id INT(2) NOT NULL,
  year INT(10) UNSIGNED NOT NULL,
  day_at VARCHAR(25) NOT NULL,
  campus VARCHAR(55) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (course_id) REFERENCES course(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE class;
