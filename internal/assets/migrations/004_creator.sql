-- +migrate Up

ALTER TABLE voting_contract_accounts
    ADD COLUMN creator_address TEXT NOT NULL,
    ADD COLUMN proposal_info_with_config JSONB NOT NULL;

-- +migrate Down

ALTER TABLE voting_contract_accounts
    DROP COLUMN creator_address,
    DROP COLUMN proposal_info_with_config;