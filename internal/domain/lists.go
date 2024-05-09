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
	ListId      string
	Name        string
	Owner       string
	Description string
	EditURL     string
	ShowURL     string
	ShareURL    string
	NewItemURL  string
	ShareCode   string
}

func ToList(list gateway.List) List {
	return List{
		ListId:      list.ListID,
		Name:        list.Name,
		Description: list.Description,
		EditURL:     fmt.Sprintf("/admin/lists/%s/edit", list.ListID),
		ShowURL:     fmt.Sprintf("/admin/lists/%s", list.ListID),
		ShareURL:    fmt.Sprintf("/shares/%s", list.ShareCode),
		NewItemURL:  fmt.Sprintf("/admin/lists/%s/items/new", list.ListID),
		ShareCode:   list.ShareCode,
	}
}

type Item struct {
	ItemId            string
	ListId            string
	Link              string
	ImageUrl          string
	Description       string
	Name              string
	Price             string
	PriceValue        string
	Quantity          int64
	PurchasedQuantity string
	Purchased         bool
	NeededQuantity    string
	ShowURL           string
	CheckoutUrl       string
	EditURL           string
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
		ItemId:            item.ListItemID,
		ListId:            item.ListID,
		Name:              item.Name,
		Link:              item.Link,
		ImageUrl:          item.ImageUrl,
		Description:       item.Description,
		Price:             moneyPrice.Display(),
		PriceValue:        formatter.Format(item.Price),
		Quantity:          item.Quantity,
		NeededQuantity:    fmt.Sprintf("%d", item.Quantity),
		PurchasedQuantity: "0",
		Purchased:         false,
		ShowURL:           fmt.Sprintf("/admin/items/%s", item.ListItemID),
		EditURL:           fmt.Sprintf("/admin/items/%s/edit", item.ListItemID),
		CheckoutUrl:       "/checkouts",
	}
}

type FindListResponse struct {
	List  List
	Items []Item
}

func (it *App) ListContainsItem(items []Item, itemId string) bool {
	for _, item := range items {
		if item.ItemId == itemId {
			return true
		}
	}

	return false
}

func (it *App) FindList(id string) (FindListResponse, error) {
	list, err := it.Queries.FindList(it.Ctx, id)

	if err != nil {
		return FindListResponse{}, err
	}

	items, err := it.Queries.FilerItemsForList(it.Ctx, list.ListID)

	if err != nil {
		return FindListResponse{}, err
	}

	return FindListResponse{
		List:  ToList(list),
		Items: utils.Map(items, ToItem),
	}, nil
}

type UpdateListParams struct {
	ListId      string
	Name        string
	Description string
}

type UpdateListResponse struct {
	List List
}

func (it *App) UpdateList(params UpdateListParams) (UpdateListResponse, error) {
	wishlist, err := it.Queries.FindList(it.Ctx, params.ListId)

	if err != nil {
		return UpdateListResponse{}, err
	}

	err = it.Queries.UpdateList(it.Ctx, gateway.UpdateListParams{
		ListID:      wishlist.ListID,
		Name:        params.Name,
		Description: params.Description,
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return UpdateListResponse{}, err
	}

	return UpdateListResponse{
		List: ToList(wishlist),
	}, nil
}

func (it *App) DeleteList(id string) error {
	list, err := it.Queries.FindList(it.Ctx, id)

	if err != nil {
		return err
	}

	err = it.Queries.DeleteList(it.Ctx, list.ListID)

	if err != nil {
		return err
	}

	return nil
}
