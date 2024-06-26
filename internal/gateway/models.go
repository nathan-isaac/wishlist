// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package gateway

import (
	"time"
)

type Checkout struct {
	CheckoutID string
	ListID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CheckoutItem struct {
	CheckoutItemID string
	CheckoutID     string
	ListItemID     string
	Quantity       int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CheckoutResponse struct {
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

type List struct {
	ListID      string
	Name        string
	Description string
	ShareCode   string
	Public      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ListItem struct {
	ListItemID  string
	ListID      string
	Link        string
	Name        string
	Description string
	ImageUrl    string
	Quantity    int64
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
