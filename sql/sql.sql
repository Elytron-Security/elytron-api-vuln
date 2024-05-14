CREATE DATABASE IF NOT EXISTS api;
USE api;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;

CREATE TABLE roles(
    id int auto_increment primary key,
    role varchar(20) not null,
    created timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(100) not null,
    role_id int not null,
    created timestamp default current_timestamp(),
    FOREIGN KEY (role_id) REFERENCES roles(id)
) ENGINE=INNODB;

INSERT INTO roles
    (role)
    values
    ("admin"),
    ("user");

INSERT INTO users
    (name, nick, email, password, role_id)
    values
    ("Administrator", "admin", "admin@api.com", "$2a$10$lJ2drtzep12SLoZ08u8nDeKiv15FyjFRXpc9qgTE8sa9oc40gN9Am", 1),
    ("User1", "Apelido1", "user1@api.com", "$2a$10$iUGvNrXjjLV46.RjCjpB8eLnQLxbLJSFSoQOMoSmIVsLXUXXfSj3.", 2),
    ("User2", "Apelido2", "user2@api.com", "$2a$10$8Mty9JsT45fQh8AOM5sL4ObbxXfkkiiVLOhfLxMAi1Wy3Eufx9UYK", 2);