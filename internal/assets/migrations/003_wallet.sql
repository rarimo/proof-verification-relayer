-- +migrate Up

CREATE TABLE IF NOT EXISTS wallets
(
    address VARCHAR(42) NOT NULL PRIMARY KEY,
    balance NUMERIC(78,0) NOT NULL,

    UNIQUE (address)
);

CREATE TABLE IF NOT EXISTS gas_limits
(
    gas_limit  NUMERIC(78,0) NOT NULL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS blocks
(
    block  BIGINT NOT NULL PRIMARY KEY
);

CREATE INDEX IF NOT EXISTS idx_address ON wallets(address);
CREATE INDEX IF NOT EXISTS idx_gas_limit ON gas_limits(gas_limit);
CREATE INDEX IF NOT EXISTS idx_block ON blocks(block);


-- +migrate Down
DROP INDEX IF EXISTS idx_address;
DROP INDEX IF EXISTS idx_gas_limit;
DROP INDEX IF EXISTS idx_block;

DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS gas_limits;
DROP TABLE IF EXISTS blocks;
