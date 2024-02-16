CREATE TABLE res (
    id INT,
    name VARCHAR(50),
    age INT,
    country VARCHAR(10),
    salary INT
);

 SELECT *FROM res;

 INSERT INTO res (id, name, age, country, salary) 
 VALUES 
    (1, 'John Doe', 30, 'USA', 5000),
    (2, 'Jane Smith', 25, 'UK', 4500),
    (3, 'Alice Johnson', 35, 'Canada', 6000),
    (4, 'Bob Brown', 40, 'Australia', 5500),
    (5, 'Eva Garcia', 28, 'Spain', 4800);
  
SELECT MIN(age)FROM res;

SELECT count(*)FROM res;

SELECT MAX(salary)FROM res;

SELECT AVG(age)FROM res;

SELECT SUM(salary)FROM res;

