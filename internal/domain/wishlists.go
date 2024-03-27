package domain

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"time"
	"wishlist/internal/gateway"
	"wishlist/internal/utils"
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
	moneyPrice := money.New(item.Price, money.USD)

	currency := money.GetCurrency(money.USD)
	formatter := money.NewFormatter(
		currency.Fraction,
		currency.Decimal,
		currency.Thousand,
		"",
		currency.Template,
	)

	return Item{
		Id:                item.ID,
		Name:              item.Name,
		Link:              item.Link,
		ImageUrl:          item.ImageUrl,
		Description:       item.Description,
		Price:             moneyPrice.Display(),
		PriceValue:        formatter.Format(item.Price),
		NeededQuantity:    fmt.Sprintf("%d", item.Quantity),
		PurchasedQuantity: "0",
		ShowURL:           fmt.Sprintf("/admin/items/%s", item.ID),
		EditURL:           fmt.Sprintf("/admin/items/%s/edit", item.ID),
	}
}

type FindWishlistResponse struct {
	Wishlist Wishlist
	Items    []Item
}

func (it *App) FindWishlist(id string) (FindWishlistResponse, error) {
	wishlist, err := it.Queries.FindWishlist(it.Ctx, id)

	if err != nil {
		return FindWishlistResponse{}, err
	}

	items, err := it.Queries.FilerItemsForWishlist(it.Ctx, id)

	if err != nil {
		return FindWishlistResponse{}, err
	}

	return FindWishlistResponse{
		Wishlist: ToWishlist(wishlist),
		Items:    utils.Map(items, ToItem),
	}, nil
}

type UpdateWishlistParams struct {
	ID          string
	Name        string
	Description string
}

type UpdateWishlistResponse struct {
	Wishlist Wishlist
}

func (it *App) UpdateWishlist(params UpdateWishlistParams) (UpdateWishlistResponse, error) {
	wishlist, err := it.Queries.FindWishlist(it.Ctx, params.ID)

	if err != nil {
		return UpdateWishlistResponse{}, err
	}

	err = it.Queries.UpdateWishlist(it.Ctx, gateway.UpdateWishlistParams{
		ID:          wishlist.ID,
		Name:        params.Name,
		Description: params.Description,
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return UpdateWishlistResponse{}, err
	}

	return UpdateWishlistResponse{
		Wishlist: ToWishlist(wishlist),
	}, nil
}

func (it *App) DeleteWishlist(id string) error {
	wishlist, err := it.Queries.FindWishlist(it.Ctx, id)

	if err != nil {
		return err
	}

	err = it.Queries.DeleteWishlist(it.Ctx, wishlist.ID)

	if err != nil {
		return err
	}

	return nil
}
