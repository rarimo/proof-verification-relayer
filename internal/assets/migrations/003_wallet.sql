-- +migrate Up

CREATE TABLE IF NOT EXISTS voting_contract_accounts
(
    voting_address VARCHAR(42) NOT NULL PRIMARY KEY,
    residual_balance NUMERIC(78,0) NOT NULL,
    gas_limit NUMERIC(78,0) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_voting_address ON voting_contract_accounts(voting_address);

CREATE TABLE IF NOT EXISTS processed_events (
    transaction_hash BYTEA NOT NULL,
    log_index BIGINT NOT NULL,
    emitted_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    block_number BIGINT NOT NULL,
    CONSTRAINT processed_events_transaction_hash_log_index_primary_key PRIMARY KEY(transaction_hash, log_index)
);

CREATE INDEX IF NOT EXISTS idx_block_number ON processed_events(block_number);

-- +migrate Down

DROP INDEX IF EXISTS idx_voting_address;
DROP INDEX IF EXISTS idx_block_number;

DROP TABLE IF EXISTS voting_contract_accounts;
DROP TABLE IF EXISTS processed_events;
