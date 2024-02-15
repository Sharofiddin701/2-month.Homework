CREATE TABLE Student (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    age INTEGER,
    gender VARCHAR(10),
    score INTEGER,
    address VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(15),
    date_of_birth DATE
);

INSERT INTO Student (id,first_name, last_name, age, gender, score, address, email, phone, date_of_birth)
VALUES
    (1,'John', 'Doe', 20, 'Male', 85, '123 Main St, Cityville', 'john.doe@example.com', '123-456-7890', '2004-05-10'),
    (2,'Jane', 'Smith', 19, 'Female', 78, '456 Elm St, Townsville', 'jane.smith@example.com', '234-567-8901', '2003-08-15'),
    (3,'Michael', 'Johnson', 21, 'Male', 92, '789 Oak St, Villageton', 'michael.johnson@example.com', '345-678-9012', '2001-12-20'),
    (4,'Emily', 'Williams', 20, 'Female', 80, '101 Pine St, Hamletville', 'emily.williams@example.com', '456-789-0123', '2002-09-25'),
    (5,'William', 'Brown', 22, 'Male', 88, '202 Cedar St, Boroughburg', 'william.brown@example.com', '567-890-1234', '2000-07-30'),
    (6,'Sophia', 'Jones', 19, 'Female', 75, '303 Maple St, Villageville', 'sophia.jones@example.com', '678-901-2345', '2003-03-05'),
    (7,'James', 'Garcia', 20, 'Male', 82, '404 Birch St, Forestburg', 'james.garcia@example.com', '789-012-3456', '2002-10-12'),
    (8,'Olivia', 'Martinez', 21, 'Female', 90, '505 Walnut St, Rivertown', 'olivia.martinez@example.com', '890-123-4567', '2001-01-18'),
    (9,'Benjamin', 'Hernandez', 20, 'Male', 86, '606 Spruce St, Hillburg', 'benjamin.hernandez@example.com', '901-234-5678', '2002-06-28'),
    (10,'Ava', 'Lopez', 19, 'Female', 79, '707 Ash St, Valleytown', 'ava.lopez@example.com', '012-345-6789', '2003-04-07');

     SELECT *FROM Student;

UPDATE Student
SET 
    first_name = 'John_New', last_name = 'Doe_New', age = 25, gender = 'Male', score = 95, 
    address = '123 New St, NewCity', email = 'john_new.doe@example.com', phone = '111-111-1111', 
    date_of_birth = '1999-01-01'
WHERE id = <5;

 SELECT *FROM Student;

 DELETE FROM Student WHERE score<80;

SELECT *FROM Student;

