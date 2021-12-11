CREATE DATABASE IF NOT EXISTS user;
USE user;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null unique,
    createAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE USER 'golangdev'@'localhost' IDENTIFIED BY 'golangdev';
GRANT ALL PRIVILEGES ON user.* TO 'golangdev'@'localhost';