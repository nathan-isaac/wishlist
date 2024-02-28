package server

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"whishlist/internal/gateway"
	"whishlist/internal/utils"
	"whishlist/internal/views"
)

func (s *Server) WishlistsShowHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return views.Render(c, views.WishlistShowView(views.ToWishlist(wishlist)))
}

func (s *Server) WishlistsEditHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return views.Render(c, views.WishlistEditView(views.ToWishlist(wishlist)))
}

func (s *Server) WishlistsUpdateHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	name := c.FormValue("name")
	description := c.FormValue("description")

	err = s.queries.UpdateWishlist(s.ctx, gateway.UpdateWishlistParams{
		ID:   wishlist.ID,
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  true,
		},
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error updating wishlist: %s", err))
	}

	return HxRedirect(c, fmt.Sprintf("/admin/wishlists/%s", wishlist.ID))
}

func (s *Server) WishlistsDeleteHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	err = s.queries.DeleteWishlist(s.ctx, wishlist.ID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error deleting wishlist: %s", err))
	}

	return HxRedirect(c, "/admin/wishlists")
}

func (s *Server) WishlistsIndexHandler(c echo.Context) error {
	wishlists, err := s.queries.ListWishlists(s.ctx)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlists: %s", err))
	}

	wishlistsView := utils.Map(wishlists, func(wishlist gateway.Wishlist) views.Wishlist {
		return views.ToWishlist(wishlist)
	})

	return views.Render(c, views.WishlistIndexView(
		views.WishlistIndex{
			NewWishlistURL: "/admin/wishlists/new",
			Wishlists:      wishlistsView,
		},
	))
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

	shareId, err := GenerateShareId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating share id: %s", err))
	}

	err = s.queries.CreateWishlist(s.ctx, gateway.CreateWishlistParams{
		ID:   id,
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  true,
		},
		ShareCode: shareId,
		Public:    false,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating wishlist: %s", err))
	}

	return c.Redirect(http.StatusFound, "/admin/wishlists")
}
