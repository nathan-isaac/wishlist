package domain

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"time"
	"wishlist/internal/gateway"
	"wishlist/internal/utils"
)

type ListIndex struct {
	NewWishlistURL string
	Lists          []List
}

type List struct {
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

func ToList(list gateway.List) List {
	return List{
		ID:          list.ID,
		Name:        list.Name,
		Description: list.Description,
		EditURL:     fmt.Sprintf("/admin/lists/%s/edit", list.ID),
		ShowURL:     fmt.Sprintf("/admin/lists/%s", list.ID),
		ShareURL:    fmt.Sprintf("/share/%s", list.ShareCode),
		NewItemURL:  fmt.Sprintf("/admin/lists/%s/items/new", list.ID),
		ShareCode:   list.ShareCode,
	}
}

func ToItem(item gateway.ListItem) Item {
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

type FindListResponse struct {
	List  List
	Items []Item
}

func (it *App) FindList(id string) (FindListResponse, error) {
	wishlist, err := it.Queries.FindList(it.Ctx, id)

	if err != nil {
		return FindListResponse{}, err
	}

	items, err := it.Queries.FilerItemsForList(it.Ctx, id)

	if err != nil {
		return FindListResponse{}, err
	}

	return FindListResponse{
		List:  ToList(wishlist),
		Items: utils.Map(items, ToItem),
	}, nil
}

type UpdateListParams struct {
	ID          string
	Name        string
	Description string
}

type UpdateListResponse struct {
	Wishlist List
}

func (it *App) UpdateWishlist(params UpdateListParams) (UpdateListResponse, error) {
	wishlist, err := it.Queries.FindList(it.Ctx, params.ID)

	if err != nil {
		return UpdateListResponse{}, err
	}

	err = it.Queries.UpdateList(it.Ctx, gateway.UpdateListParams{
		ID:          wishlist.ID,
		Name:        params.Name,
		Description: params.Description,
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return UpdateListResponse{}, err
	}

	return UpdateListResponse{
		Wishlist: ToList(wishlist),
	}, nil
}

func (it *App) DeleteWishlist(id string) error {
	list, err := it.Queries.FindList(it.Ctx, id)

	if err != nil {
		return err
	}

	err = it.Queries.DeleteList(it.Ctx, list.ID)

	if err != nil {
		return err
	}

	return nil
}
