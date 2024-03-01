-- +migrate Up
create table proofs(
    id         uuid default gen_random_uuid(),
    tx_data    text not null,
    created_at timestamp default now()
);

-- +migrate Down
