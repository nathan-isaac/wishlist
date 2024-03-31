// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: checkout.sql

package gateway

import (
	"context"
	"time"
)

const createCheckout = `-- name: CreateCheckout :exec
INSERT INTO checkout (id, list_id, created_at, updated_at)
VALUES (?, ?, ?, ?)
`

type CreateCheckoutParams struct {
	ID        string
	ListID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateCheckout(ctx context.Context, arg CreateCheckoutParams) error {
	_, err := q.db.ExecContext(ctx, createCheckout,
		arg.ID,
		arg.ListID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createCheckoutItem = `-- name: CreateCheckoutItem :exec
INSERT INTO checkout_item (id, checkout_id, list_item_id, quantity, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateCheckoutItemParams struct {
	ID         string
	CheckoutID string
	ListItemID string
	Quantity   int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (q *Queries) CreateCheckoutItem(ctx context.Context, arg CreateCheckoutItemParams) error {
	_, err := q.db.ExecContext(ctx, createCheckoutItem,
		arg.ID,
		arg.CheckoutID,
		arg.ListItemID,
		arg.Quantity,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createCheckoutResponse = `-- name: CreateCheckoutResponse :exec
INSERT INTO checkout_response (id, checkout_id, name, address_line_one, address_line_two, city, state, zip, message, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateCheckoutResponseParams struct {
	ID             string
	CheckoutID     string
	Name           string
	AddressLineOne string
	AddressLineTwo string
	City           string
	State          string
	Zip            string
	Message        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (q *Queries) CreateCheckoutResponse(ctx context.Context, arg CreateCheckoutResponseParams) error {
	_, err := q.db.ExecContext(ctx, createCheckoutResponse,
		arg.ID,
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

const filterCheckoutItems = `-- name: FilterCheckoutItems :many
SELECT checkout_item.id, checkout_item.checkout_id, checkout_item.list_item_id, checkout_item.quantity, checkout_item.created_at, checkout_item.updated_at, list_item.id, list_item.list_id, list_item.link, list_item.name, list_item.description, list_item.image_url, list_item.quantity, list_item.price, list_item.created_at, list_item.updated_at
FROM checkout_item
join list_item on checkout_item.list_item_id = list_item.id
WHERE checkout_id = ?
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
			&i.CheckoutItem.ID,
			&i.CheckoutItem.CheckoutID,
			&i.CheckoutItem.ListItemID,
			&i.CheckoutItem.Quantity,
			&i.CheckoutItem.CreatedAt,
			&i.CheckoutItem.UpdatedAt,
			&i.ListItem.ID,
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

const findCheckout = `-- name: FindCheckout :one
SELECT checkout.id, checkout.list_id, checkout.created_at, checkout.updated_at, list.id, list.name, list.description, list.share_code, list.public, list.created_at, list.updated_at
FROM checkout
JOIN list on checkout.list_id = list.id
WHERE checkout.id = ?
LIMIT 1
`

type FindCheckoutRow struct {
	Checkout Checkout
	List     List
}

func (q *Queries) FindCheckout(ctx context.Context, id string) (FindCheckoutRow, error) {
	row := q.db.QueryRowContext(ctx, findCheckout, id)
	var i FindCheckoutRow
	err := row.Scan(
		&i.Checkout.ID,
		&i.Checkout.ListID,
		&i.Checkout.CreatedAt,
		&i.Checkout.UpdatedAt,
		&i.List.ID,
		&i.List.Name,
		&i.List.Description,
		&i.List.ShareCode,
		&i.List.Public,
		&i.List.CreatedAt,
		&i.List.UpdatedAt,
	)
	return i, err
}
