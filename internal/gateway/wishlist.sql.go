// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: wishlist.sql

package gateway

import (
	"context"
	"time"
)

const createWishlist = `-- name: CreateWishlist :exec
INSERT INTO wishlist (id, name, description, share_code, public, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateWishlistParams struct {
	ID          string
	Name        string
	Description string
	ShareCode   string
	Public      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateWishlist(ctx context.Context, arg CreateWishlistParams) error {
	_, err := q.db.ExecContext(ctx, createWishlist,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ShareCode,
		arg.Public,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createWishlistItem = `-- name: CreateWishlistItem :exec
INSERT INTO wishlist_item (id, wishlist_id, name, link, image_url, description, quantity, price, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateWishlistItemParams struct {
	ID          string
	WishlistID  string
	Name        string
	Link        string
	ImageUrl    string
	Description string
	Quantity    int64
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateWishlistItem(ctx context.Context, arg CreateWishlistItemParams) error {
	_, err := q.db.ExecContext(ctx, createWishlistItem,
		arg.ID,
		arg.WishlistID,
		arg.Name,
		arg.Link,
		arg.ImageUrl,
		arg.Description,
		arg.Quantity,
		arg.Price,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE
FROM wishlist_item
WHERE id = ?
`

func (q *Queries) DeleteItem(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const deleteWishlist = `-- name: DeleteWishlist :exec
DELETE
FROM wishlist
WHERE id = ?
`

func (q *Queries) DeleteWishlist(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteWishlist, id)
	return err
}

const filerItemsForWishlist = `-- name: FilerItemsForWishlist :many
SELECT id, wishlist_id, link, name, description, image_url, quantity, price, created_at, updated_at
FROM wishlist_item
WHERE wishlist_id = ?
ORDER BY name
`

func (q *Queries) FilerItemsForWishlist(ctx context.Context, wishlistID string) ([]WishlistItem, error) {
	rows, err := q.db.QueryContext(ctx, filerItemsForWishlist, wishlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WishlistItem
	for rows.Next() {
		var i WishlistItem
		if err := rows.Scan(
			&i.ID,
			&i.WishlistID,
			&i.Link,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.Quantity,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findItem = `-- name: FindItem :one
SELECT id, wishlist_id, link, name, description, image_url, quantity, price, created_at, updated_at
FROM wishlist_item
WHERE id = ?
`

func (q *Queries) FindItem(ctx context.Context, id string) (WishlistItem, error) {
	row := q.db.QueryRowContext(ctx, findItem, id)
	var i WishlistItem
	err := row.Scan(
		&i.ID,
		&i.WishlistID,
		&i.Link,
		&i.Name,
		&i.Description,
		&i.ImageUrl,
		&i.Quantity,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findWishlist = `-- name: FindWishlist :one
SELECT id, name, description, share_code, public, created_at, updated_at
FROM wishlist
WHERE id = ?
LIMIT 1
`

func (q *Queries) FindWishlist(ctx context.Context, id string) (Wishlist, error) {
	row := q.db.QueryRowContext(ctx, findWishlist, id)
	var i Wishlist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ShareCode,
		&i.Public,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findWishlistByShareCode = `-- name: FindWishlistByShareCode :one
SELECT id, name, description, share_code, public, created_at, updated_at
FROM wishlist
WHERE share_code = ?
LIMIT 1
`

func (q *Queries) FindWishlistByShareCode(ctx context.Context, shareCode string) (Wishlist, error) {
	row := q.db.QueryRowContext(ctx, findWishlistByShareCode, shareCode)
	var i Wishlist
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ShareCode,
		&i.Public,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listWishlists = `-- name: ListWishlists :many
SELECT id, name, description, share_code, public, created_at, updated_at
FROM wishlist
ORDER BY name
`

func (q *Queries) ListWishlists(ctx context.Context) ([]Wishlist, error) {
	rows, err := q.db.QueryContext(ctx, listWishlists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Wishlist
	for rows.Next() {
		var i Wishlist
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ShareCode,
			&i.Public,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateItem = `-- name: UpdateItem :exec
UPDATE wishlist_item
set name        = ?,
    link        = ?,
    description = ?,
    image_url   = ?,
    quantity    = ?,
    price       = ?,
    updated_at  = ?
WHERE id = ?
`

type UpdateItemParams struct {
	Name        string
	Link        string
	Description string
	ImageUrl    string
	Quantity    int64
	Price       int64
	UpdatedAt   time.Time
	ID          string
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) error {
	_, err := q.db.ExecContext(ctx, updateItem,
		arg.Name,
		arg.Link,
		arg.Description,
		arg.ImageUrl,
		arg.Quantity,
		arg.Price,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateWishlist = `-- name: UpdateWishlist :exec
UPDATE wishlist
set name        = ?,
    description = ?,
    updated_at  = ?
WHERE id = ?
`

type UpdateWishlistParams struct {
	Name        string
	Description string
	UpdatedAt   time.Time
	ID          string
}

func (q *Queries) UpdateWishlist(ctx context.Context, arg UpdateWishlistParams) error {
	_, err := q.db.ExecContext(ctx, updateWishlist,
		arg.Name,
		arg.Description,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
