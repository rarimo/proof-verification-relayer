-- +migrate Up

CREATE TABLE state
(
    id         BIGSERIAL PRIMARY KEY,
    root       TEXT NOT NULL,
    block      INT  NOT NULL,
    tx_hash    TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (block, root)
);

-- +migrate Down
DROP TABLE state;