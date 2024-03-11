package domain

import (
	"fmt"
	"whishlist/internal/gateway"
)

type WishlistIndex struct {
	NewWishlistURL string
	Wishlists      []Wishlist
}

type Wishlist struct {
	ID          string
	Name        string
	Owner       string
	Description string
	EditURL     string
	ShowURL     string
	ShareURL    string
	ShareCode   string
}

type Item struct {
	Id                string
	Link              string
	ImageUrl          string
	Description       string
	Name              string
	Price             string
	PurchasedQuantity string
	NeededQuantity    string
}

func ToWishlist(wishlist gateway.Wishlist) Wishlist {
	return Wishlist{
		ID:          wishlist.ID,
		Name:        wishlist.Name,
		Description: wishlist.Description,
		EditURL:     fmt.Sprintf("/admin/wishlists/%s/edit", wishlist.ID),
		ShowURL:     fmt.Sprintf("/admin/wishlists/%s", wishlist.ID),
		ShareURL:    fmt.Sprintf("/share/%s", wishlist.ShareCode),
		ShareCode:   wishlist.ShareCode,
	}
}
