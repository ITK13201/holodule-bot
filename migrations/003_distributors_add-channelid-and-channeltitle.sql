-- +goose Up
-- +goose StatementBegin
ALTER TABLE distributors ADD COLUMN channel_id VARCHAR(128) NULL UNIQUE;
ALTER TABLE distributors ADD COLUMN channel_title VARCHAR(128) NULL UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE distributors DROP COLUMN channel_id;
ALTER TABLE distributors DROP COLUMN channel_title;
-- +goose StatementEnd