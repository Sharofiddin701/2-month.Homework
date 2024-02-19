CREATE TABLE Country(
    countryid uuid DEFAULT gen_random_uuid() Primary key,
    country_name varchar(30)
);

CREATE TABLE City(
    City_id uuid DEFAULT gen_random_uuid() Primary key,
    City_name varchar(30),
    country_id uuid References Contry(countryid)
);

INSERT INTO City (City_name, country_id) 
VALUES
('Paris', 'd7980471-0d5f-44a8-9b6f-c3d998cd3da4'),
('London', 'd7980471-0d5f-44a8-9b6f-c3d998cd3da4'),
('Beijing', '0e89287a-f98c-4a2a-bf4a-273015210429'),
('Moscow', 'd7980471-0d5f-44a8-9b6f-c3d998cd3da4'),
('Rome', 'd7980471-0d5f-44a8-9b6f-c3d998cd3da4'),
('Madrid', 'd7980471-0d5f-44a8-9b6f-c3d998cd3da4'),
('Tokyo', '0e89287a-f98c-4a2a-bf4a-273015210429'),
('Seoul', '0e89287a-f98c-4a2a-bf4a-273015210429'),
('Sydney', '933b50b2-e3bf-4c3f-808e-4c0e0daf087b'),
('Mumbai', '485f52cf-0afc-4400-9ead-29a1a77061ae');


SELECT City.City_name, Country.country_name FROM City JOIN Country ON City.country_id = Country.countryid;

SELECT AVG(length(City_name)) From city;

SELECT COUNT(City_name) FROM City;

SELECT Country.country_name from Country JOIN CITY  WHERE (ON COUNT(City.country_id = Country.countryid)>3);

SELECT *FROM CITY WHERE City_name like'A%' ;

ALTER TABLE City ADD COLUMN population INT;

UPDATE City

SET population = CASE
    WHEN City_name = 'Paris' THEN 3000000
    WHEN City_name = 'London' THEN 700000
    WHEN City_name = 'Beijing' THEN 8000000
    WHEN City_name = 'Moscow' THEN 2000000
    WHEN City_name = 'Rome' THEN 2000000
    WHEN City_name = 'Madrid' THEN 1000000
    WHEN City_name = 'Tokyo' THEN 20000000
    WHEN City_name = 'Seoul' THEN 10000000
    WHEN City_name = 'Sydney' THEN 3000000
    WHEN City_name = 'Mumbai' THEN 1000000
END;


SELECT SUM(population) FROM CITY;

UPDATE City  SET population=3000000 WHERE City_name='Paris';

DELETE FROM City WHERE population<100000;

TIME ADD

ALTER TABLE City ADD COLUMN TIME_ADD TIME;

UPDATE CITY SET City_name='PARIS1',TIME_ADD=current_timestamp WHERE City_id_id='d7980471-0d5f-44a8-9b6f-c3d998cd3da4';

UPDATE CITY SET population=55555555,TIME_ADD=current_timestamp WHERE City_name='Rome';



