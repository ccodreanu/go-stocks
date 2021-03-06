CREATE TABLE IF NOT EXISTS symbols (
    symbol VARCHAR(16) PRIMARY KEY,
    currency VARCHAR(3) NOT NULL
);

CREATE TABLE IF NOT EXISTS values (
    symbol VARCHAR(16) NOT NULL,
    ts timestamp without time zone NOT NULL,
    value REAL NOT NULL,
    PRIMARY KEY(symbol, ts)
);
