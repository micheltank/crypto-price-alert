CREATE TABLE IF NOT EXISTS "alert"
(
    id        SERIAL PRIMARY KEY,
    email     TEXT    NOT NULL,
    price     NUMERIC NOT NULL,
    coin      TEXT    NOT NULL,
    direction TEXT    NOT NULL
);