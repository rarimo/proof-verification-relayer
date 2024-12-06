-- +migrate Up

ALTER TABLE state
    DROP COLUMN created_at;
ALTER TABLE state
    ADD COLUMN timestamp BIGINT NOT NULL default 0;

-- +migrate Down

ALTER TABLE state
    DROP COLUMN timestamp;
ALTER TABLE state
    ADD COLUMN created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP;