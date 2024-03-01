-- +migrate Up
create table proofs(
    id                 uuid default gen_random_uuid(),
    voting_session     text not null,
    document_nullifier text not null,
    created_at         timestamp default now(),
    unique (voting_session, document_nullifier)
);

-- +migrate Down
drop table proofs;