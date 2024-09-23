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

************************************************************************************************

CREATE INDEX idx_name ON users (name);

DROP INDEX idx_name ON users;

                                                   ***
inner join:
SELECT users.name, orders.order_date
FROM users
INNER JOIN orders ON users.id = orders.user_id;

left join:
SELECT users.name, orders.order_date
FROM users
LEFT JOIN orders ON users.id = orders.user_id;

right join:
SELECT users.name, orders.order_date
FROM users
RIGHT JOIN orders ON users.id = orders.user_id;

full join:
SELECT users.name, orders.order_date
FROM users
FULL OUTER JOIN orders ON users.id = orders.user_id;

cross join:
SELECT users.name, orders.order_date
FROM users
CROSS JOIN orders;

transaction:
START TRANSACTION;

INSERT INTO users (name, phone_number) VALUES ('Alice', '09991234567');
UPDATE users SET phone_number = '0987654321' WHERE name = 'Bob';

COMMIT;  


ROLLBACK; 

aggregate functions:
SELECT COUNT(*) AS total_users FROM users;
SELECT AVG(order_amount) AS average_order FROM orders;
SELECT MAX(order_amount) AS max_order, MIN(order_amount) AS min_order FROM orders;
SELECT SUM(order_amount) AS total_revenue FROM orders;

group by and having:
SELECT user_id, COUNT(*) AS total_orders
FROM orders
GROUP BY user_id
HAVING total_orders > 5;