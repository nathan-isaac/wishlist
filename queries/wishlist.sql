-- name: FindWishlist :one
SELECT *
FROM wishlist
WHERE id = ? LIMIT 1;

-- name: FindWishlistByShareCode :one
SELECT *
FROM wishlist
WHERE share_code = ? LIMIT 1;

-- name: ListWishlists :many
SELECT *
FROM wishlist
ORDER BY name;

-- name: CreateWishlist :exec
INSERT INTO wishlist (id, name, description, share_code, public)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateWishlist :exec
UPDATE wishlist
set name = ?,
    description  = ?
WHERE id = ?;

-- name: DeleteWishlist :exec
DELETE
FROM wishlist
WHERE id = ?;

-- name: CreateWishlistItem :exec
INSERT INTO wishlist_item (id, wishlist_id, name, link, image_url, description, quantity, price)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: FilerItemsForWishlist :many
SELECT *
FROM wishlist_item
WHERE wishlist_id = ?
ORDER BY name;
