USE sql_store;

SELECT customer_id, first_name, last_name, birth_date, points, points+10
FROM customers
WHERE points>500 OR MONTH (birth_date) = '4'
ORDER BY points desc;

SELECT DISTINCT state
FROM customers;

SELECT 
name, 
unit_price,
unit_price *1.1 AS 'new price'
FROM products;

SELECT*
FROM orders
WHERE YEAR (order_date) >= '2018';

SELECT*, unit_price*quantity AS total_price
FROM order_items
WHERE order_id = 6 OR unit_price*quantity > 30
ORDER BY unit_price*quantity desc;

SELECT*
FROM customers
WHERE address LIKE '%trail%' OR 
      address LIKE '%avenue%' OR 
      phone LIKE '%9';

SELECT*
FROM customers
WHERE first_name REGEXP 'elka|ambur';

SELECT*
FROM customers
WHERE last_name REGEXP 'ey$|on$';

SELECT*
FROM customers
WHERE last_name REGEXP '^my|se';

SELECT*
FROM customers
WHERE last_name REGEXP 'b[r,u]';

SELECT*
FROM orders
WHERE shipper_id IS NULL;

SELECT *, quantity * unit_price AS total_price
FROM order_items
WHERE order_id = '2'
ORDER BY quantity*unit_price DESC;

SELECT*
FROM customers
ORDER BY points DESC
LIMIT 4;

SELECT order_id, oi.product_id, name, quantity, oi.unit_price
FROM order_items oi
JOIN products p
  ON oi.product_id = p.product_id;
  
  SELECT p.product_id, p.name, oi.quantity
  FROM products p
 LEFT JOIN order_items oi
    ON oi.product_id = p.product_id AND oi.product_id != 3 
 WHERE p.product_id IN (2, 3, 10, 11, 12);
    
SELECT o.order_date, o.order_id, c.first_name, s.name AS shipper, os.name AS status
FROM orders o
LEFT JOIN customers c
  ON c.customer_id = o.customer_id
LEFT JOIN shippers s
  ON o.shipper_id = s.shipper_id
LEFT JOIN order_statuses os
  ON os.order_status_id = o.status
  ORDER BY order_id;

	SELECT 
	customer_id, 
	first_name, 
	points,
	'Bronze' AS type
	FROM customers	
	WHERE points < 2000 
UNION
	SELECT 
	customer_id, 
	first_name, 
	points,
	'Silver' AS type
	FROM customers	
	WHERE points BETWEEN 2000 AND 3000
UNION
	SELECT 
	customer_id, 
	first_name, 
	points,
	'Gold' AS type
	FROM customers	
	WHERE points > 3000;

INSERT INTO products (product_id, name, quantity_in_stock, unit_price)
VALUES (DEFAULT, 'product1', 1, '1.1'),
	   (DEFAULT, 'product2', 2, '2.2'),
       (DEFAULT, 'product3', 3, '3.3');
       
INSERT INTO orders (order_id, customer_id, order_date, status)
VALUES (DEFAULT, 2, '2019-02-02', 1);

INSERT INTO order_items (order_id, product_id, quantity, unit_price)
VALUES (LAST_INSERT_ID (), 7, 2, 3.2),
       (LAST_INSERT_ID (), 1, 4, 4.2);


SELECT*
FROM customers
WHERE birth_date >'1990-01-01';

UPDATE customers
SET points = points+50
WHERE birth_date >'1990-01-01';
 
UPDATE orders
SET comments = 'gold'
WHERE customer_id IN
(SELECT customer_id
FROM customers 
WHERE points > 3000);

SELECT*
FROM products p 
CROSS JOIN customers c;

SELECT*
FROM customers c
LEFT JOIN orders o
USING (customer_id)
ORDER BY c.customer_id;

SELECT 			o.order_id, 
				p.product_id,
                o.status
                p.q
                uantity_in_stock 
  FROM order_items oi
 LEFT JOIN products p
    ON oi.product_id = p.product_id
    JOIN orders o
    ON o.order_id = oi.order_id
ORDER BY p.product_id;

SELECT*
FROM employees e
JOIN employees m
   ON e.reports_to = m.employee_id;
   
SELECT*
FROM products p 
CROSS JOINS customers c







