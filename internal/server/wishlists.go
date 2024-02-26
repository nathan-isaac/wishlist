package server

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"whishlist/internal/gateway"
	"whishlist/internal/views"
)

func (s *Server) WishlistsShowHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return c.JSON(http.StatusOK, wishlist)
}

func (s *Server) WishlistsNewHandler(c echo.Context) error {
	return views.Render(c, views.WishlistNewView())
}

func (s *Server) WishlistsPostHandler(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	id, err := GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating id: %s", err))
	}

	err = s.queries.CreateWishlist(s.ctx, gateway.CreateWishlistParams{
		ID:   id,
		Name: name,
		Description: sql.NullString{
			String: description,
		},
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating wishlist: %s", err))
	}

	return c.Redirect(http.StatusFound, "/admin/wishlists")
}
