-- +migrate Up

ALTER TABLE voting_contract_accounts
    ADD COLUMN creator_address TEXT NOT NULL,
    ADD COLUMN proposal_info_with_config JSONB NOT NULL;

CREATE INDEX IF NOT EXISTS idx_creator_address ON voting_contract_accounts(creator_address);
CREATE INDEX IF NOT EXISTS idx_proposal_info_with_config ON voting_contract_accounts USING GIN(proposal_info_with_config jsonb_path_ops);

-- +migrate Down

DROP INDEX IF EXISTS idx_creator_address;
DROP INDEX IF EXISTS idx_proposal_info_with_config;

ALTER TABLE voting_contract_accounts
    DROP COLUMN creator_address,
    DROP COLUMN proposal_info_with_config;