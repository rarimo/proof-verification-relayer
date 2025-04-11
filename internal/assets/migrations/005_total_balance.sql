-- +migrate Up

DROP INDEX IF EXISTS idx_proposal_info_with_config;

ALTER TABLE voting_contract_accounts
    ADD COLUMN total_balance NUMERIC(78,0) NOT NULL,
    ADD COLUMN min_age INT NOT NULL,
    ADD COLUMN max_age INT NOT NULL,
    ADD COLUMN start_timestamp BIGINT NOT NULL,
    ADD COLUMN end_timestamp BIGINT NOT NULL,
    ADD COLUMN votes_count INT NOT NULL;


ALTER TABLE voting_contract_accounts
    RENAME COLUMN proposal_info_with_config TO parsed_whitelist_data_with_metadata;

CREATE INDEX IF NOT EXISTS idx_parsed_whitelist_data_with_metadata ON voting_contract_accounts USING GIN(parsed_whitelist_data_with_metadata jsonb_path_ops);
CREATE INDEX IF NOT EXISTS idx_min_age ON voting_contract_accounts(min_age);
CREATE INDEX IF NOT EXISTS idx_max_age ON voting_contract_accounts(max_age);

-- +migrate Down

DROP INDEX IF EXISTS idx_parsed_whitelist_data_with_metadata;
DROP INDEX IF EXISTS idx_min_age;
DROP INDEX IF EXISTS idx_max_age;

ALTER TABLE voting_contract_accounts
    DROP COLUMN total_balance,
    DROP COLUMN min_age,
    DROP COLUMN max_age,
    DROP COLUMN start_timestamp,
    DROP COLUMN end_timestamp,
    DROP COLUMN votes_count;


ALTER TABLE voting_contract_accounts
    RENAME COLUMN parsed_whitelist_data_with_metadata TO proposal_info_with_config;

CREATE INDEX IF NOT EXISTS idx_proposal_info_with_config ON voting_contract_accounts USING GIN(proposal_info_with_config jsonb_path_ops);


