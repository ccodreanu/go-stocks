CREATE TABLE IF NOT EXISTS symbols (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(16) UNIQUE,
    currency VARCHAR(3) NOT NULL
);

CREATE TABLE IF NOT EXISTS historical_values (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(16) NOT NULL,
    ts timestamp without time zone NOT NULL,
    value REAL NOT NULL
);
