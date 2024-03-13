package domain

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"strings"
	"wishlist/internal/gateway"
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
	NewItemURL  string
	ShareCode   string
}

type Item struct {
	Id                string
	Link              string
	ImageUrl          string
	Description       string
	Name              string
	Price             string
	PriceValue        string
	PurchasedQuantity string
	NeededQuantity    string
	ShowURL           string
	EditURL           string
}

func ToWishlist(wishlist gateway.Wishlist) Wishlist {
	return Wishlist{
		ID:          wishlist.ID,
		Name:        wishlist.Name,
		Description: wishlist.Description,
		EditURL:     fmt.Sprintf("/admin/wishlists/%s/edit", wishlist.ID),
		ShowURL:     fmt.Sprintf("/admin/wishlists/%s", wishlist.ID),
		ShareURL:    fmt.Sprintf("/share/%s", wishlist.ShareCode),
		NewItemURL:  fmt.Sprintf("/admin/wishlists/%s/items/new", wishlist.ID),
		ShareCode:   wishlist.ShareCode,
	}
}

func ToItem(item gateway.WishlistItem) Item {
	moneyPrice := money.New(item.Price, "USD")

	return Item{
		Id:                item.ID,
		Name:              item.Name,
		Link:              item.Link,
		ImageUrl:          item.ImageUrl,
		Description:       item.Description,
		Price:             moneyPrice.Display(),
		PriceValue:        strings.Replace(moneyPrice.Display(), "$", "", -1),
		NeededQuantity:    fmt.Sprintf("%d", item.Quantity),
		PurchasedQuantity: "0",
		ShowURL:           fmt.Sprintf("/admin/items/%s", item.ID),
		EditURL:           fmt.Sprintf("/admin/items/%s/edit", item.ID),
	}
}
