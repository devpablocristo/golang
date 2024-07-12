CREATE DATABASE IF NOT EXISTS meli_items DEFAULT COLLATE utf8mb4_general_ci DEFAULT CHARSET utf8mb4; USE
    meli_items;
CREATE TABLE item(
    id INT NOT NULL AUTO_INCREMENT,
    CODE VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    STATUS ENUM
        ('ACTIVE', 'INACTIVE') NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL,
        PRIMARY KEY(id),
        UNIQUE KEY CODE(CODE)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT COLLATE = utf8mb4_general_ci DEFAULT CHARSET utf8mb4;
INSERT INTO item(
    CODE,
    title,
    description,
    price,
    stock,
STATUS
    ,
    created_at,
    updated_at
)
VALUES(
    'SAM27324354',
    'Tablet Samsung Galaxy Tab S7',
    'Galaxy Tab S7 with S Pen SM-t733 12.4 pulgadas y 64GB de memoria RAM',
    150000.76,
    3,
    'ACTIVE',
    '2020-05-10 04:20:33',
    '2020-05-10 05:30:00'
),(
    'APP58439385',
    'iPhone 13 Pro Max',
    'Nuevo iPhone 13 Pro Max con pantalla OLED de 6.7 pulgadas y cámara triple',
    180000.43,
    5,
    'ACTIVE',
    '2022-03-28 10:30:00',
    '2022-04-01 16:45:00'
),(
    'DLX98993240',
    'Laptop Dell XPS 13',
    'Laptop Dell XPS 13 con procesador Intel Core i7 y 16GB de RAM',
    120000.23,
    2,
    'ACTIVE',
    '2021-10-15 09:00:00',
    '2021-10-15 10:30:00'
),(
    'SAM48502349',
    'Smartphone Samsung Galaxy S21',
    'Samsung Galaxy S21 con pantalla AMOLED de 6.2 pulgadas y 128GB de RAM',
    90000,
    8,
    'ACTIVE',
    '2021-02-10 14:20:00',
    '2021-02-11 09:30:00'
),(
    'APP73940923',
    'iPad Air',
    'iPad Air de 10.9 pulgadas con pantalla Liquid Retina y Touch ID',
    70000,
    4,
    'INACTIVE',
    '2020-09-05 18:00:00',
    '2020-09-08 11:15:00'
);