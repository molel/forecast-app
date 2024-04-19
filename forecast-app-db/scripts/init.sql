CREATE TABLE users
(
    username VARCHAR PRIMARY KEY,
    password CHAR(128) NOT NULL
);

CREATE TABLE measurements_units
(
    name VARCHAR PRIMARY KEY
);


CREATE TABLE time_series
(
    id       BIGSERIAL PRIMARY KEY,
    username varchar NOT NULL REFERENCES users (username),
    unit     VARCHAR NOT NULL REFERENCES measurements_units (name),
    name     varchar NOT NULL
);

CREATE TABLE records
(
    series_id BIGINT REFERENCES time_series (id),
    ts        TIMESTAMPTZ,
    value     FLOAT NOT NULL,
    PRIMARY KEY (series_id, ts)
);