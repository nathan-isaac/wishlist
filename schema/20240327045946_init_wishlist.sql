-- +goose Up
-- +goose StatementBegin
CREATE TABLE wishlist
(
    id          TEXT PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        NOT NULL,
    share_code  TEXT UNIQUE NOT NULL,
    public      BOOLEAN     NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL
);
CREATE TABLE wishlist_item
(
    id          TEXT PRIMARY KEY,
    wishlist_id TEXT      NOT NULL,
    link        TEXT      NOT NULL,
    name        TEXT      NOT NULL,
    description TEXT      NOT NULL,
    image_url   TEXT      NOT NULL,
    quantity    INTEGER   NOT NULL,
    price       INTEGER   NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    FOREIGN KEY (wishlist_id) REFERENCES wishlist (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wishlist;
DROP TABLE wishlist_item;
-- +goose StatementEnd
