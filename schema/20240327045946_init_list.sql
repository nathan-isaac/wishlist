-- +goose Up
-- +goose StatementBegin
CREATE TABLE list
(
    list_id          TEXT PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        NOT NULL,
    share_code  TEXT UNIQUE NOT NULL,
    public      BOOLEAN     NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL
);
CREATE TABLE list_item
(
    list_item_id          TEXT PRIMARY KEY,
    list_id TEXT      NOT NULL,
    link        TEXT      NOT NULL,
    name        TEXT      NOT NULL,
    description TEXT      NOT NULL,
    image_url   TEXT      NOT NULL,
    quantity    INTEGER   NOT NULL,
    price       INTEGER   NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    FOREIGN KEY (list_id) REFERENCES list (list_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE list;
DROP TABLE list_item;
-- +goose StatementEnd
