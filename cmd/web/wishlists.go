package web

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WishlistController struct {
	db *sqlx.DB
}

func NewWishlists(db *sqlx.DB) *WishlistController {
	return &WishlistController{
		db: db,
	}
}

func (w WishlistController) WishlistShowHandler(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, fmt.Sprintf("showing wishlist %s", id))
}
