-- +goose Up
-- +goose StatementBegin
CREATE TABLE checkout
(
    id          TEXT PRIMARY KEY,
    list_id TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    FOREIGN KEY (list_id) REFERENCES list (id)
);
create table checkout_response
(
    id               TEXT PRIMARY KEY,
    checkout_id      TEXT      NOT NULL,
    name             TEXT      NOT NULL,
    address_line_one TEXT      NOT NULL,
    address_line_two TEXT      NOT NULL,
    city             TEXT      NOT NULL,
    state            TEXT      NOT NULL,
    zip              TEXT      NOT NULL,
    message          TEXT      NOT NULL,
    created_at       TIMESTAMP NOT NULL,
    updated_at       TIMESTAMP NOT NULL,
    FOREIGN KEY (checkout_id) REFERENCES checkout (id)
);
create table checkout_item
(
    id               TEXT PRIMARY KEY,
    checkout_id      TEXT      NOT NULL,
    list_item_id TEXT      NOT NULL,
    quantity         INTEGER   NOT NULL,
    created_at       TIMESTAMP NOT NULL,
    updated_at       TIMESTAMP NOT NULL,
    FOREIGN KEY (checkout_id) REFERENCES checkout (id),
    FOREIGN KEY (list_item_id) REFERENCES list_item (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE checkout;
DROP TABLE checkout_response;
DROP TABLE checkout_item;
-- +goose StatementEnd
