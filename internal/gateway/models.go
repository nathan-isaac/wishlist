// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package gateway

import (
	"database/sql"
	"time"
)

type Wishlist struct {
	ID          string
	Name        string
	Description sql.NullString
}

type WishlistItem struct {
	ID          string
	WishlistID  sql.NullString
	Link        string
	Description sql.NullString
	WantedCount int64
}

type WishlistPurchase struct {
	ID             string
	WishlistItemID sql.NullString
	BoughtAt       time.Time
}
