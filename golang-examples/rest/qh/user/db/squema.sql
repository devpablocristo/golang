CREATE DATABASE max_inventory;
USE max_inventory;

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE products (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  created_by INT NOT NULL,
  PRIMARY KEY (id)
  FOREIGN (created_by) REFERENCES users(id)
);

CREATE TABLE roles (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE user_roles (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  role_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN (user_id) REFERENCES users(id),
  FOREIGN (role_id) REFERENCES roles(id)
);
