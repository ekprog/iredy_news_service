-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS news
(
    id          bigint unique,
    title       VARCHAR(255)      not null default 0,
    image       VARCHAR(255)      not null default 0,
    type        VARCHAR(10)       not null ,

    created_at  timestamp(0) NOT NULL DEFAULT now(),
    updated_at  timestamp(0) NOT NULL DEFAULT now(),
    deleted_at  timestamp(0)        DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "news";
-- +goose StatementEnd