CREATE TABLE todos (
  todo_id BIGINT NOT NULL PRIMARY KEY,
  activity_group_id CHAR(36),
  title VARCHAR(255),
  priority INT,
  is_active BOOLEAN,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (activity_group_id) REFERENCES activities(activity_id)
);
