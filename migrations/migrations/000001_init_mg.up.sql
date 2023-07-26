CREATE TABLE users
(
    id BIGINT PRIMARY KEY,
    balance INTEGER DEFAULT NULL
)

CREATE TABLE transactions
(
    id BIGINT PRIMARY KEY,
    user_from INTEGER NOT NULL,
    user_to INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL
)

INSERT INTO users (id, balance) VALUES(1,trunc(random()))
INSERT INTO users (id, balance) VALUES(2,trunc(random()))