CREATE DATABASE my_database;

USE my_database;

CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, phone_number) VALUES ('abootaleb', '09123336789');

SELECT * FROM users;
/* 
filter:
*/
SELECT * FROM users WHERE name = 'Hadi';

UPDATE users SET phone_number = '0987654321' WHERE name = 'Ali';

DELETE FROM users WHERE name = 'shahrzad';

ALTER TABLE users ADD COLUMN email VARCHAR(255) UNIQUE;

ALTER TABLE users MODIFY COLUMN phone_number VARCHAR(15);
ALTER TABLE users CHANGE COLUMN name full_name VARCHAR(255);

ALTER TABLE users DROP COLUMN email;

DROP TABLE users;