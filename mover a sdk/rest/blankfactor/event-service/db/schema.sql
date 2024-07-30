CREATE DATABASE IF NOT EXISTS events_service DEFAULT COLLATE utf8mb4_general_ci DEFAULT CHARSET utf8mb4;

USE events_service;

CREATE TABLE events (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    PRIMARY KEY(id)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT COLLATE = utf8mb4_general_ci DEFAULT CHARSET utf8mb4;

INSERT INTO events(start_time, end_time)
VALUES(
    '2023-04-13 09:00:00',
    '2023-04-13 11:00:00'
),(
    '2023-04-13 10:00:00',
    '2023-04-13 12:00:00'
),(
    '2023-04-13 14:00:00',
    '2023-04-13 16:00:00'
),(
    '2023-04-14 09:00:00',
    '2023-04-14 11:00:00'
),(
    '2023-04-15 11:00:00',
    '2023-04-15 13:00:00'
),(
    '2023-04-16 13:00:00',
    '2023-04-16 15:00:00'
);