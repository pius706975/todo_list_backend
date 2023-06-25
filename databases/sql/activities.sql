-- Active: 1687592118175@@localhost@3306@todo_list
CREATE TABLE activities (
    activity_id CHAR(36) NOT NULL DEFAULT (UUID()) PRIMARY KEY,
    title VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);