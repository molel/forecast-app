CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE,
    password CHAR(128) NOT NULL
);

CREATE TABLE measurement_units
(
    id   SMALLSERIAL PRIMARY KEY,
    name VARCHAR UNIQUE
);

CREATE TABLE time_series
(
    user_id          INT      NOT NULL REFERENCES users (id),
    unit_id          SMALLINT NOT NULL REFERENCES measurement_units (id),
    name             VARCHAR  NOT NULL,
    period           SMALLINT NOT NULL,
    prediction_start SMALLINT NOT NULL,
    id               BIGSERIAL UNIQUE,
    PRIMARY KEY (user_id, name)
);

CREATE TABLE records
(
    series_id BIGINT REFERENCES time_series (id),
    ts        BIGINT,
    value     REAL NOT NULL,
    PRIMARY KEY (series_id, ts)
);
