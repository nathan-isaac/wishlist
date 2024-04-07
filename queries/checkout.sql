-- name: CreateCheckout :exec
INSERT INTO checkout (id, list_id, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- name: UpdateCheckout :exec
UPDATE checkout
SET updated_at = ?
WHERE id = ?;

-- name: CreateCheckoutItem :exec
INSERT INTO checkout_item (id, checkout_id, list_item_id, quantity, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateCheckoutResponse :exec
INSERT INTO checkout_response (id, checkout_id, name, address_line_one, address_line_two, city, state, zip, message, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateCheckoutItem :exec
UPDATE checkout_item
SET quantity = ?, updated_at = ?
WHERE id = ?;

-- name: UpdateCheckoutResponse :exec
UPDATE checkout_response
SET name = ?, address_line_one = ?, address_line_two = ?, city = ?, state = ?, zip = ?, message = ?, updated_at = ?
WHERE id = ?;

-- name: FindCheckout :one
SELECT sqlc.embed(checkout), sqlc.embed(list)
FROM checkout
JOIN list on checkout.list_id = list.id
WHERE checkout.id = ?
LIMIT 1;

-- name: FindCheckoutResponse :one
SELECT *
FROM checkout_response
WHERE checkout_id = ?
LIMIT 1;

-- name: FilterCheckoutItems :many
SELECT sqlc.embed(checkout_item), sqlc.embed(list_item)
FROM checkout_item
join list_item on checkout_item.list_item_id = list_item.id
WHERE checkout_id = ?;
