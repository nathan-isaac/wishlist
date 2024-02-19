package web

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"whishlist/internal/gateway"
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

	wishlist, err := w.queries.GetWishlist(w.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return c.JSON(http.StatusOK, wishlist)
}
