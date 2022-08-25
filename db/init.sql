SET SCHEMA 'public';

-- TABLES
CREATE TABLE orders
(
    id        SERIAL PRIMARY KEY,
    order_uid varchar(64) NOT NULL,
    data      jsonb       NOT NULL
);

-- PROCEDURES
CREATE PROCEDURE insert_order(
    _order_uid varchar(64),
    _data jsonb
)
    LANGUAGE sql AS
$BODY$
    INSERT INTO orders (order_uid, data)
    VALUES (_order_uid, _data);
$BODY$;

-- FUNCTION
CREATE FUNCTION get_all_orders()
    RETURNS TABLE (data jsonb)
    LANGUAGE sql AS
$BODY$
    SELECT data
    FROM orders;
$BODY$;

CREATE FUNCTION get_order_by_id(_id varchar(64))
    RETURNS TABLE (data jsonb)
    LANGUAGE sql AS
$BODY$
    SELECT data
    FROM orders
    WHERE order_uid = _id;
$BODY$;