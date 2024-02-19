-- name: GetWishlist :one
SELECT *
FROM wishlist
WHERE id = ? LIMIT 1;

-- name: ListWishlists :many
SELECT *
FROM wishlist
ORDER BY name;

-- name: CreateWishlist :exec
INSERT INTO wishlist (id, name, description)
VALUES (?, ?, ?);

-- name: UpdateWishlist :exec
UPDATE wishlist
set name = ?,
    description  = ?
WHERE id = ?;

-- name: DeleteWishlist :exec
DELETE
FROM wishlist
WHERE id = ?;
