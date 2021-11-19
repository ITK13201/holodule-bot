-- +goose Up
-- +goose StatementBegin
CREATE TABLE videos (
    id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
    distributor_id INTEGER NOT NULL,
    url VARCHAR(256) NOT NULL DEFAULT '',
    datetime DATETIME NOT NULL,
    image_url VARCHAR(256) NOT NULL DEFAULT '',
    notified_at DATETIME NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_videos_distributor_id
        FOREIGN KEY (distributor_id)
        REFERENCES distributors (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS videos;
-- +goose StatementEnd