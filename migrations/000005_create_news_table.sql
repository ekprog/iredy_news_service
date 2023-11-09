-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS news (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) not null,
        image VARCHAR(255),
        type VARCHAR(10) not null,
        is_active BOOLEAN,
        created_at timestamp(0) NOT NULL DEFAULT now (),
        updated_at timestamp(0) NOT NULL DEFAULT now (),
        deleted_at timestamp(0) DEFAULT NULL
    );

CREATE TABLE
    IF NOT EXISTS news_details (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) not null,
        image VARCHAR(255),
        type VARCHAR(10) NOT NULL,

        news_id INTEGER REFERENCES news(id) not null,

        swipe_delay INTEGER not null,
        is_active BOOLEAN,

        created_at timestamp(0) NOT NULL DEFAULT now (),
        updated_at timestamp(0) NOT NULL DEFAULT now (),
        deleted_at timestamp(0) DEFAULT NULL
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "news";
DROP TABLE IF EXISTS "news_details";

-- +goose StatementEnd