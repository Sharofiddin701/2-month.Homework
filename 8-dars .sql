2-masala.

create table order_products (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(20),
    price integer,
    order_id uuid references orders(id)
)

create table orders (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id integer,
    amount integer
)


insert into orders (external_id, amount) values(456789,100000),
(123123,120000),
(182131,21333),
(123123,2749713),
(24234,139595);

insert into order_products (name, price, order_id) 
 VALUES('Apple1',1000,'42c730c6-ee71-4902-8867-500ab305f62a'),
 ('samsung',2000,'cacb8cec-2992-438f-97d9-5d4da5e796b7'),
 ('mi',3000,'8e6a0f13-07bc-4643-8abf-fcf1e2fe380f'),
 ('vivo',1232,'0f6234e2-3ce0-4902-8a71-d3ffae623aa1'),
 ('awdaw',1221,'59d20fb2-2825-4d92-a87c-ee3694883e1d');



create or replace function totalprice()
returns integer
LANGUAGE plpgsql
as
$$
begin
return(
select  sum(order_products.price*orders.amount)
from order_products join orders
On order_products.order_id=orders.id
);
end;
$$;
 
drop function totalprice;

select * from totalprice();


CREATE OR REPLACE FUNCTION totalprice()
RETURNS INTEGER
LANGUAGE plpgsql
AS
$$
BEGIN
  RETURN (
    SELECT SUM(order_products.price * orders.amount) as totalprice
    FROM order_products JOIN orders 
    ON order_products.order_id = orders.id
  );
END;
$$;



create or replace function totalprice()
returns float
LANGUAGE plpgsql
as
$$
begin
return  (
 SELECT SUM(order_products.price * orders.amount) as totalprice
    FROM order_products JOIN orders 
    on order_products.order_id = orders.id
);
end;
$$;

drop function totalprice();

SELECT * from totalprice() ;

3-masala.

CREATE INDEX index_order_id
ON order_products(order_id);

4-masala.

CREATE TABLE order_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(20),
    price INTEGER,
    order_id UUID REFERENCES orders(id)
);

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id INTEGER,
    amount INTEGER
);

INSERT INTO orders (external_id, amount) VALUES (456789, 100000),
                                                (123123, 120000),
                                                (182131, 21333),
                                                (123123, 2749713),
                                                (24234, 139595);

INSERT INTO order_products (name, price, order_id) VALUES ('Apple1', 1000, '42c730c6-ee71-4902-8867-500ab305f62a'),
                                                           ('samsung', 2000, 'cacb8cec-2992-438f-97d9-5d4da5e796b7'),
                                                           ('mi', 3000, '8e6a0f13-07bc-4643-8abf-fcf1e2fe380f'),
                                                           ('vivo', 1232, '0f6234e2-3ce0-4902-8a71-d3ffae623aa1'),
                                                           ('awdaw', 1221, '59d20fb2-2825-4d92-a87c-ee3694883e1d');

5-masala.

CREATE INDEX index_order_id_name
ON order_products(name, order_id);

7-masala.

CREATE OR REPLACE FUNCTION gettotal(order_idd UUID)
RETURNS INT
LANGUAGE plpgsql
AS
$$
DECLARE 
    Totalprice INT;
BEGIN
    SELECT SUM(order_products.price * orders.amount)
    INTO Totalprice
    FROM order_products
    JOIN orders ON order_products.order_id = orders.id
    WHERE orders.id = order_idd;

    RETURN Totalprice;
END;
$$;

SELECT gettotal('42c730c6-ee71-4902-8867-500ab305f62a');


