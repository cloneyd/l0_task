SET SCHEMA 'public';

-- TABLES
CREATE TABLE deliveries
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR NOT NULL,
    phone   VARCHAR NOT NULL,
    zip     VARCHAR NOT NULL,
    city    VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    region  VARCHAR NOT NULL,
    email   VARCHAR NOT NULL
);

CREATE TABLE payments
(
    id            SERIAL PRIMARY KEY,
    transaction   VARCHAR NOT NULL,
    request_id    VARCHAR NOT NULL,
    currency      VARCHAR NOT NULL,
    provider      VARCHAR NOT NULL,
    amount        INT     NOT NULL,
    payment_dt    INT     NOT NULL,
    bank          VARCHAR NOT NULL,
    delivery_cost INT     NOT NULL,
    goods_total   INT     NOT NULL,
    custom_fee    INT     NOT NULL
);

CREATE TABLE orders
(
    id                 SERIAL PRIMARY KEY,
    order_uid          VARCHAR UNIQUE NOT NULL,
    track_number       VARCHAR        NOT NULL,
    entry              VARCHAR        NOT NULL,
    delivery_id        INT            NOT NULL,
    payment_id         INT            NOT NULL,
    locale             VARCHAR        NOT NULL,
    internal_signature VARCHAR        NOT NULL,
    customer_id        VARCHAR        NOT NULL,
    delivery_service   VARCHAR        NOT NULL,
    shardkey           VARCHAR        NOT NULL,
    sm_id              int            NOT NULL,
    date_created       VARCHAR        NOT NULL,
    oof_shard          VARCHAR        NOT NULL,
    CONSTRAINT fk_delivery_id
        FOREIGN KEY (delivery_id)
            REFERENCES deliveries (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT fk_payment_id
        FOREIGN KEY (payment_id)
            REFERENCES payments (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);

CREATE UNIQUE INDEX idx_order_uid
    ON orders (order_uid);

CREATE TABLE items
(
    id           SERIAL PRIMARY KEY,
    chrt_id      INT     NOT NULL,
    track_number VARCHAR NOT NULL,
    price        INT     NOT NULL,
    rid          VARCHAR NOT NULL,
    sale         INT     NOT NULL,
    size         VARCHAR NOT NULL,
    total_price  INT     NOT NULL,
    nm_id        INT     NOT NULL,
    brand        VARCHAR NOT NULL,
    status       INT     NOT NULL
);

CREATE TABLE items_to_orders
(
    item_id  INT NOT NULL,
    order_id INT NOT NULL,
    CONSTRAINT fk_item_id
        FOREIGN KEY (item_id)
            REFERENCES items (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
    CONSTRAINT fk_order_id
        FOREIGN KEY (order_id)
            REFERENCES orders (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);

CREATE UNIQUE INDEX idx_item_order
    ON items_to_orders (item_id, order_id);


-- FUNCTIONS
CREATE FUNCTION insert_order(order_uid varchar,
                             track_number varchar,
                             entry varchar,
                             delivery_id INT,
                             payment_id INT,
                             locale varchar,
                             internal_signature varchar,
                             customer_id varchar,
                             delivery_service varchar,
                             shardkey varchar,
                             sm_id int,
                             date_created varchar,
                             oof_shard varchar)
    RETURNS TABLE (id int)
    LANGUAGE SQL AS $BODY$
INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
    VALUES (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        RETURNING id;
$BODY$;

CREATE FUNCTION get_all_orders()
    RETURNS TABLE (id int,
        order_uid varchar,
        track_number varchar,
        entry varchar,
        delivery_id INT,
        payment_id INT,
        locale varchar,
        internal_signature varchar,
        customer_id varchar,
        delivery_service varchar,
        shardkey varchar,
        sm_id varchar,
        date_created varchar,
        oof_shard varchar)
    LANGUAGE SQL AS
$BODY$
SELECT *
FROM orders;
$BODY$;

CREATE FUNCTION get_order_by_id(uid varchar)
    RETURNS TABLE (id int,
        order_uid varchar,
        track_number varchar,
        entry varchar,
        delivery_id INT,
        payment_id INT,
        locale varchar,
        internal_signature varchar,
        customer_id varchar,
        delivery_service varchar,
        shardkey varchar,
        sm_id int,
        date_created varchar,
        oof_shard varchar)
    LANGUAGE SQL AS
$BODY$
SELECT *
FROM orders
WHERE order_uid = uid;
$BODY$;

CREATE FUNCTION insert_delivery(name varchar,
                                phone varchar,
                                zip varchar,
                                city varchar,
                                address varchar,
                                region varchar,
                                email varchar)
    RETURNS TABLE (id int)
    LANGUAGE SQL AS
$BODY$
INSERT INTO deliveries (name, phone, zip, city, address, region, email)
VALUES (name, phone, zip, city, address, region, email)
RETURNING id;
$BODY$;

CREATE FUNCTION insert_payment("transaction" varchar,
                               request_id varchar,
                               currency varchar,
                               provider varchar,
                               amount int,
                               payment_dt int,
                               bank varchar,
                               delivery_cost int,
                               goods_total int,
                               custom_fee int)
    RETURNS TABLE (id int)
    LANGUAGE SQL AS
$BODY$
INSERT INTO payments ("transaction", request_id, currency, provider, amount, payment_dt, bank, delivery_cost,
                      goods_total, custom_fee)
VALUES ("transaction", request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total,
        custom_fee)
RETURNING id;
$BODY$;

CREATE FUNCTION insert_item(chrt_id int,
                            track_number varchar,
                            price int,
                            rid varchar,
                            sale int,
                            "size" varchar,
                            total_price int,
                            nm_id int,
                            brand varchar,
                            status int)
    RETURNS TABLE (id int)
    LANGUAGE SQL AS
$BODY$
INSERT INTO items (chrt_id, track_number, price, rid, sale, "size", total_price, nm_id, brand, status)
VALUES (chrt_id, track_number, price, rid, sale, "size", total_price, nm_id, brand, status)
RETURNING id;
$BODY$;

CREATE FUNCTION get_items_by_order_id(orderId int)
    RETURNS TABLE ()
    LANGUAGE SQL AS
$BODY$
SELECT *
FROM orders
WHERE id = orderId;
$BODY$;

CREATE PROCEDURE insert_items_to_orders(item_id int, order_id int)
    LANGUAGE SQL AS
$BODY$
INSERT INTO items_to_orders (item_id, order_id)
VALUES (item_id, order_id);
$BODY$;