-- +migrate Up
create table proofs(
    id                 uuid default gen_random_uuid(),
    registration       text not null,
    document_nullifier text not null,
    created_at         timestamp default now(),
    unique (registration, document_nullifier)
);

-- +migrate Down
drop table proofs;