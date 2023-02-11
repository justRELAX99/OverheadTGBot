-- +goose Up
CREATE TABLE IF NOT EXISTS message(
    id INTEGER PRIMARY KEY,
    text TEXT NOT NULL,
    date integer not null,
    status TEXT not null
);

-- +goose Down
DROP TABLE IF EXISTS message;