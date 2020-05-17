-- +goose Up
CREATE TABLE IF NOT EXISTS course_class (
	course_id INT(10) UNSIGNED NOT NULL,
	class_id INT(10) UNSIGNED NOT NULL,
	FOREIGN KEY (course_id) REFERENCES course(id),
	FOREIGN KEY (class_id) REFERENCES class(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE course_class;
