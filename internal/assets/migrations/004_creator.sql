-- +migrate Up

ALTER TABLE voting_contract_accounts
    ADD COLUMN creator_address TEXT NOT NULL;

-- +migrate Down

ALTER TABLE voting_contract_accounts
    DROP COLUMN creator_address;
