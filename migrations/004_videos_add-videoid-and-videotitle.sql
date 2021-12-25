-- +goose Up
-- +goose StatementBegin
ALTER TABLE videos ADD COLUMN video_id VARCHAR(128) NULL UNIQUE;
ALTER TABLE videos ADD COLUMN video_title VARCHAR(128) NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE videos DROP COLUMN video_id;
ALTER TABLE videos DROP COLUMN video_title;
-- +goose StatementEnd