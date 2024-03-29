// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package gateway

import (
	"time"
)

type Checkout struct {
	ID         string
	WishlistID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CheckoutItem struct {
	ID             string
	CheckoutID     string
	WishlistItemID string
	Quantity       int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CheckoutResponse struct {
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

type Wishlist struct {
	ID          string
	Name        string
	Description string
	ShareCode   string
	Public      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type WishlistItem struct {
	ID          string
	WishlistID  string
	Link        string
	Name        string
	Description string
	ImageUrl    string
	Quantity    int64
	Price       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
