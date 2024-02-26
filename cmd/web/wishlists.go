package web

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"whishlist/internal/gateway"
	"whishlist/internal/views"
)

type WishlistController struct {
	ctx     context.Context
	queries *gateway.Queries
}

func NewWishlists(ctx context.Context, queries *gateway.Queries) *WishlistController {
	return &WishlistController{
		ctx:     ctx,
		queries: queries,
	}
}

func (w WishlistController) WishlistShowHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := w.queries.FindWishlist(w.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return c.JSON(http.StatusOK, wishlist)
}

func (w WishlistController) ShareShowHandler(c echo.Context) error {
	code := c.Param("code")

	wishlist, err := w.queries.FindWishlistByShareCode(w.ctx, sql.NullString{String: code, Valid: true})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	items, err := w.queries.ListWishlistsItemsForWishlist(w.ctx, wishlist.ID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist items: %s", err))
	}

	wishlistItems := make([]views.Item, len(items))

	for i, item := range items {
		wishlistItems[i] = views.Item{
			Id:                item.ID,
			Name:              item.Name,
			Link:              item.Link,
			ImageUrl:          item.ImageUrl.String,
			Description:       item.Description.String,
			Price:             fmt.Sprintf("%d", item.Price),
			NeededQuantity:    fmt.Sprintf("%d", item.Quantity),
			PurchasedQuantity: "0",
		}
	}

	share := views.Share{
		Id:   wishlist.ID,
		Code: code,
		Wishlist: views.Wishlist{
			ID:          wishlist.ID,
			Name:        wishlist.Name,
			Owner:       "Owner",
			Description: wishlist.Description.String,
		},
		Items: wishlistItems,
	}

	return views.Render(c, views.ShareView(share))
}
