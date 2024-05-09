-- name: CreateCheckout :exec
INSERT INTO checkout (checkout_id, list_id, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- name: UpdateCheckout :exec
UPDATE checkout
SET updated_at = ?
WHERE checkout_id = ?;

-- name: CreateCheckoutItem :exec
INSERT INTO checkout_item (checkout_item_id, checkout_id, list_item_id, quantity, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateCheckoutItem :exec
UPDATE checkout_item
SET quantity   = ?,
    updated_at = ?
WHERE checkout_item_id = ?;

-- name: DeleteCheckoutItem :exec
DELETE FROM checkout_item
where checkout_item_id = ?;

-- name: CreateCheckoutResponse :exec
INSERT INTO checkout_response (checkout_response_id, checkout_id, name, address_line_one, address_line_two, city, state, zip, message,
                               created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateCheckoutResponse :exec
UPDATE checkout_response
SET name             = ?,
    address_line_one = ?,
    address_line_two = ?,
    city             = ?,
    state            = ?,
    zip              = ?,
    message          = ?,
    updated_at       = ?
WHERE checkout_response_id = ?;

-- name: FindCheckout :one
SELECT sqlc.embed(checkout), sqlc.embed(list)
FROM checkout
         JOIN list on checkout.list_id = list.list_id
WHERE checkout.checkout_id = ?
LIMIT 1;

-- name: FindCheckoutResponse :one
SELECT *
FROM checkout_response
WHERE checkout_id = ?
LIMIT 1;

-- name: FilterCheckoutItems :many
SELECT sqlc.embed(checkout_item), sqlc.embed(list_item)
FROM checkout_item
         join list_item on checkout_item.list_item_id = list_item.list_item_id
WHERE checkout_item.checkout_id = ?;

-- name: FindCheckoutItem :one
SELECT sqlc.embed(checkout_item), sqlc.embed(list_item)
FROM checkout_item
         join list_item on checkout_item.list_item_id = list_item.list_item_id
WHERE checkout_item.checkout_item_id = ?;

-- name: FindCheckoutItemByItemId :one
SELECT sqlc.embed(checkout_item), sqlc.embed(list_item)
FROM checkout_item
         join list_item on checkout_item.list_item_id = list_item.list_item_id
WHERE checkout_item.checkout_id = ?
  and checkout_item.list_item_id = ?
order by list_item.name;

-- name: FilterCheckoutItemsByListId :many
SELECT sqlc.embed(checkout_item), sqlc.embed(list_item), sqlc.embed(checkout)
from checkout_item
         join checkout on checkout_item.checkout_id = checkout.checkout_id
         join list_item on checkout_item.list_item_id = list_item.list_item_id
where list_item.list_id = ?;
