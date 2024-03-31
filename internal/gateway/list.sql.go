// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: list.sql

package gateway

import (
	"context"
	"time"
)

const createList = `-- name: CreateList :exec
INSERT INTO list (id, name, description, share_code, public, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateListParams struct {
	ID          string
	Name        string
	Description string
	ShareCode   string
	Public      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) error {
	_, err := q.db.ExecContext(ctx, createList,
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

const createListItem = `-- name: CreateListItem :exec
INSERT INTO list_item (id, list_id, name, link, image_url, description, quantity, price, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateListItemParams struct {
	ID          string
	ListID      string
	Name        string
	Link        string
	ImageUrl    string
	Description string
	Quantity    int64
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateListItem(ctx context.Context, arg CreateListItemParams) error {
	_, err := q.db.ExecContext(ctx, createListItem,
		arg.ID,
		arg.ListID,
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
FROM list_item
WHERE id = ?
`

func (q *Queries) DeleteItem(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const deleteList = `-- name: DeleteList :exec
DELETE
FROM list
WHERE id = ?
`

func (q *Queries) DeleteList(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteList, id)
	return err
}

const filerItemsForList = `-- name: FilerItemsForList :many
SELECT id, list_id, link, name, description, image_url, quantity, price, created_at, updated_at
FROM list_item
WHERE list_id = ?
ORDER BY name
`

func (q *Queries) FilerItemsForList(ctx context.Context, listID string) ([]ListItem, error) {
	rows, err := q.db.QueryContext(ctx, filerItemsForList, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListItem
	for rows.Next() {
		var i ListItem
		if err := rows.Scan(
			&i.ID,
			&i.ListID,
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

const filterLists = `-- name: FilterLists :many
SELECT id, name, description, share_code, public, created_at, updated_at
FROM list
ORDER BY name
`

func (q *Queries) FilterLists(ctx context.Context) ([]List, error) {
	rows, err := q.db.QueryContext(ctx, filterLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []List
	for rows.Next() {
		var i List
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

const findItem = `-- name: FindItem :one
SELECT id, list_id, link, name, description, image_url, quantity, price, created_at, updated_at
FROM list_item
WHERE id = ?
`

func (q *Queries) FindItem(ctx context.Context, id string) (ListItem, error) {
	row := q.db.QueryRowContext(ctx, findItem, id)
	var i ListItem
	err := row.Scan(
		&i.ID,
		&i.ListID,
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

const findList = `-- name: FindList :one
SELECT id, name, description, share_code, public, created_at, updated_at
FROM list
WHERE id = ?
LIMIT 1
`

func (q *Queries) FindList(ctx context.Context, id string) (List, error) {
	row := q.db.QueryRowContext(ctx, findList, id)
	var i List
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

const findListByShareCode = `-- name: FindListByShareCode :one
SELECT id, name, description, share_code, public, created_at, updated_at
FROM list
WHERE share_code = ?
LIMIT 1
`

func (q *Queries) FindListByShareCode(ctx context.Context, shareCode string) (List, error) {
	row := q.db.QueryRowContext(ctx, findListByShareCode, shareCode)
	var i List
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

const updateItem = `-- name: UpdateItem :exec
UPDATE list_item
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

const updateList = `-- name: UpdateList :exec
UPDATE list
set name        = ?,
    description = ?,
    updated_at  = ?
WHERE id = ?
`

type UpdateListParams struct {
	Name        string
	Description string
	UpdatedAt   time.Time
	ID          string
}

func (q *Queries) UpdateList(ctx context.Context, arg UpdateListParams) error {
	_, err := q.db.ExecContext(ctx, updateList,
		arg.Name,
		arg.Description,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}