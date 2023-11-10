CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    username VARCHAR NOT NULL
)

CREATE TABLE wallets (
    wallet_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_date timestamp NOT NULL DEFAULT NOW(),
    balance numeric(10,4) DEFAULT 0,
    currency VARCHAR,
    user_id uuid,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(user_id)
)

CREATE TABLE transactions (
    tran_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    tran_type VARCHAR,
    amount numeric(10,4),
    tran_date timestamp NOT NULL DEFAULT NOW(),
    wallet_id uuid,
    CONSTRAINT fk_wallet
        FOREIGN KEY(wallet_id)
            REFERENCES wallets(wallet_id)
)