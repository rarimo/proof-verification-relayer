-- +migrate Up

CREATE TABLE state
(
    id         uuid PRIMARY KEY            DEFAULT gen_random_uuid(),
    root       TEXT NOT NULL,
    block      INT  NOT NULL,
    tx_hash    TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (root, tx_hash)
);

-- +migrate Down
DROP TABLE state;