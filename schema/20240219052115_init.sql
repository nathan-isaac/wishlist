-- +goose Up
-- +goose StatementBegin
CREATE TABLE wishlist
(
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT,
    share_code  TEXT UNIQUE
);
CREATE TABLE wishlist_item
(
    id           TEXT PRIMARY KEY,
    wishlist_id  TEXT,
    link         TEXT    NOT NULL,
    description  TEXT,
    wanted_count INTEGER NOT NULL,
    FOREIGN KEY (wishlist_id) REFERENCES wishlist (id)
);
CREATE TABLE wishlist_purchase
(
    id               TEXT PRIMARY KEY,
    wishlist_item_id TEXT,
    bought_at        DATETIME NOT NULL,
    FOREIGN KEY (wishlist_item_id) REFERENCES wishlist_item (id)
);
CREATE TABLE wishlist_address
(
    id          TEXT PRIMARY KEY,
    wishlist_id TEXT,
    address     TEXT NOT NULL,
    FOREIGN KEY (wishlist_id) REFERENCES wishlist (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wishlist_purchase;
DROP TABLE wishlist_item;
DROP TABLE wishlist;
DROP TABLE wishlist_address;
-- +goose StatementEnd
