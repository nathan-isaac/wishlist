-- name: CreateCheckout :exec
INSERT INTO checkout (id, list_id, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- name: CreateCheckoutItem :exec
INSERT INTO checkout_item (id, checkout_id, list_item_id, quantity, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateCheckoutResponse :exec
INSERT INTO checkout_response (id, checkout_id, name, address_line_one, address_line_two, city, state, zip, message, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
