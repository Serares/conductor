-- +goose Up
CREATE TABLE peers (
    id UUID PRIMARY KEY,
    ip VARCHAR NOT NULL,
    transactionHash VARCHAR NOT NULL,
    fileHash VARCHAR NOT NULL,
    pubKey VARCHAR NOT NULL
);
-- +goose Down
DROP TABLE peers;