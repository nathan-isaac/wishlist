-- name: FindList :one
SELECT *
FROM list
WHERE list_id = ?
LIMIT 1;

-- name: FindListByShareCode :one
SELECT *
FROM list
WHERE share_code = ?
LIMIT 1;

-- name: FilterLists :many
SELECT *
FROM list
ORDER BY name;

-- name: CreateList :exec
INSERT INTO list (list_id, name, description, share_code, public, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateList :exec
UPDATE list
set name        = ?,
    description = ?,
    updated_at  = ?
WHERE list_id = ?;

-- name: DeleteList :exec
DELETE
FROM list
WHERE list_id = ?;

-- name: CreateListItem :exec
INSERT INTO list_item (list_item_id, list_id, name, link, image_url, description, quantity, price, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: FilerItemsForList :many
SELECT *
FROM list_item
WHERE list_id = ?
ORDER BY name;

-- name: FindItem :one
SELECT *
FROM list_item
WHERE list_item_id = ?;

-- name: UpdateItem :exec
UPDATE list_item
set name        = ?,
    link        = ?,
    description = ?,
    image_url   = ?,
    quantity    = ?,
    price       = ?,
    updated_at  = ?
WHERE list_item_id = ?;

-- name: DeleteItem :exec
DELETE
FROM list_item
WHERE list_item_id = ?;
