package web

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
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

	wishlist, err := w.queries.GetWishlist(w.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return c.JSON(http.StatusOK, wishlist)
}

func (w WishlistController) ShareShowHandler(c echo.Context) error {
	code := c.Param("code")

	return render(c, views.ShareView(code))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
