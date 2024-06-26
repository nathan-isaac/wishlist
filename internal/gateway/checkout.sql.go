// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: checkout.sql

package gateway

import (
	"context"
	"time"
)

const createCheckout = `-- name: CreateCheckout :exec
INSERT INTO checkout (checkout_id, list_id, created_at, updated_at)
VALUES (?, ?, ?, ?)
`

type CreateCheckoutParams struct {
	CheckoutID string
	ListID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (q *Queries) CreateCheckout(ctx context.Context, arg CreateCheckoutParams) error {
	_, err := q.db.ExecContext(ctx, createCheckout,
		arg.CheckoutID,
		arg.ListID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createCheckoutItem = `-- name: CreateCheckoutItem :exec
INSERT INTO checkout_item (checkout_item_id, checkout_id, list_item_id, quantity, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateCheckoutItemParams struct {
	CheckoutItemID string
	CheckoutID     string
	ListItemID     string
	Quantity       int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (q *Queries) CreateCheckoutItem(ctx context.Context, arg CreateCheckoutItemParams) error {
	_, err := q.db.ExecContext(ctx, createCheckoutItem,
		arg.CheckoutItemID,
		arg.CheckoutID,
		arg.ListItemID,
		arg.Quantity,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createCheckoutResponse = `-- name: CreateCheckoutResponse :exec
INSERT INTO checkout_response (checkout_response_id, checkout_id, name, address_line_one, address_line_two, city, state, zip, message,
                               created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateCheckoutResponseParams struct {
	CheckoutResponseID string
	CheckoutID         string
	Name               string
	AddressLineOne     string
	AddressLineTwo     string
	City               string
	State              string
	Zip                string
	Message            string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (q *Queries) CreateCheckoutResponse(ctx context.Context, arg CreateCheckoutResponseParams) error {
	_, err := q.db.ExecContext(ctx, createCheckoutResponse,
		arg.CheckoutResponseID,
		arg.CheckoutID,
		arg.Name,
		arg.AddressLineOne,
		arg.AddressLineTwo,
		arg.City,
		arg.State,
		arg.Zip,
		arg.Message,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteCheckoutItem = `-- name: DeleteCheckoutItem :exec
DELETE FROM checkout_item
where checkout_item_id = ?
`

func (q *Queries) DeleteCheckoutItem(ctx context.Context, checkoutItemID string) error {
	_, err := q.db.ExecContext(ctx, deleteCheckoutItem, checkoutItemID)
	return err
}

const filterCheckoutItems = `-- name: FilterCheckoutItems :many
SELECT checkout_item.checkout_item_id, checkout_item.checkout_id, checkout_item.list_item_id, checkout_item.quantity, checkout_item.created_at, checkout_item.updated_at, list_item.list_item_id, list_item.list_id, list_item.link, list_item.name, list_item.description, list_item.image_url, list_item.quantity, list_item.price, list_item.created_at, list_item.updated_at
FROM checkout_item
         join list_item on checkout_item.list_item_id = list_item.list_item_id
WHERE checkout_item.checkout_id = ?
`

type FilterCheckoutItemsRow struct {
	CheckoutItem CheckoutItem
	ListItem     ListItem
}

func (q *Queries) FilterCheckoutItems(ctx context.Context, checkoutID string) ([]FilterCheckoutItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, filterCheckoutItems, checkoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FilterCheckoutItemsRow
	for rows.Next() {
		var i FilterCheckoutItemsRow
		if err := rows.Scan(
			&i.CheckoutItem.CheckoutItemID,
			&i.CheckoutItem.CheckoutID,
			&i.CheckoutItem.ListItemID,
			&i.CheckoutItem.Quantity,
			&i.CheckoutItem.CreatedAt,
			&i.CheckoutItem.UpdatedAt,
			&i.ListItem.ListItemID,
			&i.ListItem.ListID,
			&i.ListItem.Link,
			&i.ListItem.Name,
			&i.ListItem.Description,
			&i.ListItem.ImageUrl,
			&i.ListItem.Quantity,
			&i.ListItem.Price,
			&i.ListItem.CreatedAt,
			&i.ListItem.UpdatedAt,
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

const filterCheckoutItemsByListId = `-- name: FilterCheckoutItemsByListId :many
SELECT checkout_item.checkout_item_id, checkout_item.checkout_id, checkout_item.list_item_id, checkout_item.quantity, checkout_item.created_at, checkout_item.updated_at, list_item.list_item_id, list_item.list_id, list_item.link, list_item.name, list_item.description, list_item.image_url, list_item.quantity, list_item.price, list_item.created_at, list_item.updated_at, checkout.checkout_id, checkout.list_id, checkout.created_at, checkout.updated_at
from checkout_item
         join checkout on checkout_item.checkout_id = checkout.checkout_id
         join list_item on checkout_item.list_item_id = list_item.list_item_id
where list_item.list_id = ?
`

type FilterCheckoutItemsByListIdRow struct {
	CheckoutItem CheckoutItem
	ListItem     ListItem
	Checkout     Checkout
}

func (q *Queries) FilterCheckoutItemsByListId(ctx context.Context, listID string) ([]FilterCheckoutItemsByListIdRow, error) {
	rows, err := q.db.QueryContext(ctx, filterCheckoutItemsByListId, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FilterCheckoutItemsByListIdRow
	for rows.Next() {
		var i FilterCheckoutItemsByListIdRow
		if err := rows.Scan(
			&i.CheckoutItem.CheckoutItemID,
			&i.CheckoutItem.CheckoutID,
			&i.CheckoutItem.ListItemID,
			&i.CheckoutItem.Quantity,
			&i.CheckoutItem.CreatedAt,
			&i.CheckoutItem.UpdatedAt,
			&i.ListItem.ListItemID,
			&i.ListItem.ListID,
			&i.ListItem.Link,
			&i.ListItem.Name,
			&i.ListItem.Description,
			&i.ListItem.ImageUrl,
			&i.ListItem.Quantity,
			&i.ListItem.Price,
			&i.ListItem.CreatedAt,
			&i.ListItem.UpdatedAt,
			&i.Checkout.CheckoutID,
			&i.Checkout.ListID,
			&i.Checkout.CreatedAt,
			&i.Checkout.UpdatedAt,
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

const findCheckout = `-- name: FindCheckout :one
SELECT checkout.checkout_id, checkout.list_id, checkout.created_at, checkout.updated_at, list.list_id, list.name, list.description, list.share_code, list.public, list.created_at, list.updated_at
FROM checkout
         JOIN list on checkout.list_id = list.list_id
WHERE checkout.checkout_id = ?
LIMIT 1
`

type FindCheckoutRow struct {
	Checkout Checkout
	List     List
}

func (q *Queries) FindCheckout(ctx context.Context, checkoutID string) (FindCheckoutRow, error) {
	row := q.db.QueryRowContext(ctx, findCheckout, checkoutID)
	var i FindCheckoutRow
	err := row.Scan(
		&i.Checkout.CheckoutID,
		&i.Checkout.ListID,
		&i.Checkout.CreatedAt,
		&i.Checkout.UpdatedAt,
		&i.List.ListID,
		&i.List.Name,
		&i.List.Description,
		&i.List.ShareCode,
		&i.List.Public,
		&i.List.CreatedAt,
		&i.List.UpdatedAt,
	)
	return i, err
}

const findCheckoutItem = `-- name: FindCheckoutItem :one
SELECT checkout_item.checkout_item_id, checkout_item.checkout_id, checkout_item.list_item_id, checkout_item.quantity, checkout_item.created_at, checkout_item.updated_at, list_item.list_item_id, list_item.list_id, list_item.link, list_item.name, list_item.description, list_item.image_url, list_item.quantity, list_item.price, list_item.created_at, list_item.updated_at
FROM checkout_item
         join list_item on checkout_item.list_item_id = list_item.list_item_id
WHERE checkout_item.checkout_item_id = ?
`

type FindCheckoutItemRow struct {
	CheckoutItem CheckoutItem
	ListItem     ListItem
}

func (q *Queries) FindCheckoutItem(ctx context.Context, checkoutItemID string) (FindCheckoutItemRow, error) {
	row := q.db.QueryRowContext(ctx, findCheckoutItem, checkoutItemID)
	var i FindCheckoutItemRow
	err := row.Scan(
		&i.CheckoutItem.CheckoutItemID,
		&i.CheckoutItem.CheckoutID,
		&i.CheckoutItem.ListItemID,
		&i.CheckoutItem.Quantity,
		&i.CheckoutItem.CreatedAt,
		&i.CheckoutItem.UpdatedAt,
		&i.ListItem.ListItemID,
		&i.ListItem.ListID,
		&i.ListItem.Link,
		&i.ListItem.Name,
		&i.ListItem.Description,
		&i.ListItem.ImageUrl,
		&i.ListItem.Quantity,
		&i.ListItem.Price,
		&i.ListItem.CreatedAt,
		&i.ListItem.UpdatedAt,
	)
	return i, err
}

const findCheckoutItemByItemId = `-- name: FindCheckoutItemByItemId :one
SELECT checkout_item.checkout_item_id, checkout_item.checkout_id, checkout_item.list_item_id, checkout_item.quantity, checkout_item.created_at, checkout_item.updated_at, list_item.list_item_id, list_item.list_id, list_item.link, list_item.name, list_item.description, list_item.image_url, list_item.quantity, list_item.price, list_item.created_at, list_item.updated_at
FROM checkout_item
         join list_item on checkout_item.list_item_id = list_item.list_item_id
WHERE checkout_item.checkout_id = ?
  and checkout_item.list_item_id = ?
order by list_item.name
`

type FindCheckoutItemByItemIdParams struct {
	CheckoutID string
	ListItemID string
}

type FindCheckoutItemByItemIdRow struct {
	CheckoutItem CheckoutItem
	ListItem     ListItem
}

func (q *Queries) FindCheckoutItemByItemId(ctx context.Context, arg FindCheckoutItemByItemIdParams) (FindCheckoutItemByItemIdRow, error) {
	row := q.db.QueryRowContext(ctx, findCheckoutItemByItemId, arg.CheckoutID, arg.ListItemID)
	var i FindCheckoutItemByItemIdRow
	err := row.Scan(
		&i.CheckoutItem.CheckoutItemID,
		&i.CheckoutItem.CheckoutID,
		&i.CheckoutItem.ListItemID,
		&i.CheckoutItem.Quantity,
		&i.CheckoutItem.CreatedAt,
		&i.CheckoutItem.UpdatedAt,
		&i.ListItem.ListItemID,
		&i.ListItem.ListID,
		&i.ListItem.Link,
		&i.ListItem.Name,
		&i.ListItem.Description,
		&i.ListItem.ImageUrl,
		&i.ListItem.Quantity,
		&i.ListItem.Price,
		&i.ListItem.CreatedAt,
		&i.ListItem.UpdatedAt,
	)
	return i, err
}

const findCheckoutResponse = `-- name: FindCheckoutResponse :one
SELECT checkout_response_id, checkout_id, name, address_line_one, address_line_two, city, state, zip, message, created_at, updated_at
FROM checkout_response
WHERE checkout_id = ?
LIMIT 1
`

func (q *Queries) FindCheckoutResponse(ctx context.Context, checkoutID string) (CheckoutResponse, error) {
	row := q.db.QueryRowContext(ctx, findCheckoutResponse, checkoutID)
	var i CheckoutResponse
	err := row.Scan(
		&i.CheckoutResponseID,
		&i.CheckoutID,
		&i.Name,
		&i.AddressLineOne,
		&i.AddressLineTwo,
		&i.City,
		&i.State,
		&i.Zip,
		&i.Message,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCheckout = `-- name: UpdateCheckout :exec
UPDATE checkout
SET updated_at = ?
WHERE checkout_id = ?
`

type UpdateCheckoutParams struct {
	UpdatedAt  time.Time
	CheckoutID string
}

func (q *Queries) UpdateCheckout(ctx context.Context, arg UpdateCheckoutParams) error {
	_, err := q.db.ExecContext(ctx, updateCheckout, arg.UpdatedAt, arg.CheckoutID)
	return err
}

const updateCheckoutItem = `-- name: UpdateCheckoutItem :exec
UPDATE checkout_item
SET quantity   = ?,
    updated_at = ?
WHERE checkout_item_id = ?
`

type UpdateCheckoutItemParams struct {
	Quantity       int64
	UpdatedAt      time.Time
	CheckoutItemID string
}

func (q *Queries) UpdateCheckoutItem(ctx context.Context, arg UpdateCheckoutItemParams) error {
	_, err := q.db.ExecContext(ctx, updateCheckoutItem, arg.Quantity, arg.UpdatedAt, arg.CheckoutItemID)
	return err
}

const updateCheckoutResponse = `-- name: UpdateCheckoutResponse :exec
UPDATE checkout_response
SET name             = ?,
    address_line_one = ?,
    address_line_two = ?,
    city             = ?,
    state            = ?,
    zip              = ?,
    message          = ?,
    updated_at       = ?
WHERE checkout_response_id = ?
`

type UpdateCheckoutResponseParams struct {
	Name               string
	AddressLineOne     string
	AddressLineTwo     string
	City               string
	State              string
	Zip                string
	Message            string
	UpdatedAt          time.Time
	CheckoutResponseID string
}

func (q *Queries) UpdateCheckoutResponse(ctx context.Context, arg UpdateCheckoutResponseParams) error {
	_, err := q.db.ExecContext(ctx, updateCheckoutResponse,
		arg.Name,
		arg.AddressLineOne,
		arg.AddressLineTwo,
		arg.City,
		arg.State,
		arg.Zip,
		arg.Message,
		arg.UpdatedAt,
		arg.CheckoutResponseID,
	)
	return err
}
