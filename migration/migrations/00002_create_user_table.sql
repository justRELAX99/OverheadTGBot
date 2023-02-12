-- +goose Up
CREATE TABLE IF NOT EXISTS user(
    id INTEGER PRIMARY KEY,
    telegram_id integer NOT NULL,
    user_name TEXT not null,
    role TEXT not null
);

-- +goose Down
DROP TABLE IF EXISTS user;