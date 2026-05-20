CREATE DATABASE socialplus;
\c socialplus;
DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id int auto_increment PRIMARY KEY,
    nome VARCHAR(60) NOT NULL,
    nick VARCHAR(60) NOT NULL UNIQUE,
    email VARCHAR(60) NOT NULL UNIQUE,
    pass VARCHAR(20) NOT NULL,
    createdAt TIMESTAMP DEFAULT current_timestap()
); 

