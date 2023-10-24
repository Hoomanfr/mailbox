CREATE TABLE IF NOT EXISTS mailboxes
(
    id VARCHAR(200) NOT NULL,
    user_id VARCHAR(200),
    email VARCHAR(200),
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id)
);