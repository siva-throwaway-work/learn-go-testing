CREATE TABLE IF NOT EXISTS customers (id int, name varchar(255), email varchar(255));
CREATE TABLE IF NOT EXISTS preferences (id int, cust_id int, subscribed bool);
CREATE TABLE IF NOT EXISTS orders (id int, amount int);

INSERT INTO customers(id, name, email) VALUES (1, 'John', 'john@gmail.com');
INSERT INTO customers(id, name, email) VALUES (2, 'Paul', 'paul@gmail.com');
INSERT INTO preferences(id, cust_id, subscribed) VALUES (1, 1, true);