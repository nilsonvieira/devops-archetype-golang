CREATE DATABASE goapi;
USE goapi;

CREATE TABLE users (
    id INT AUTO_INCREMENT,
    name VARCHAR(100),
    email VARCHAR(100),
    PRIMARY KEY (id)
);

INSERT INTO users (name, email) VALUES 
('Nilson Vieira', 'nilsonrsvieira@gmail.com'),
('Fabiola Garreto', 'fabiolagarreto@gmail.com'),
('Nilson Vieira 2', 'nilsonrsvieira@outlok.com');
