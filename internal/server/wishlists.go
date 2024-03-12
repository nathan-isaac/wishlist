package server

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"wishlist/internal/domain"
	"wishlist/internal/gateway"
	"wishlist/internal/utils"
	"wishlist/internal/views"
)

func (s *Server) WishlistsShowHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	items, err := s.queries.FilerItemsForWishlist(s.ctx, wishlist.ID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist items: %s", err))
	}

	wishlistItems := make([]domain.Item, len(items))

	for i, item := range items {
		wishlistItems[i] = domain.ToItem(item)
	}

	return views.Render(c, views.WishlistShowView(domain.ToWishlist(wishlist), wishlistItems))
}

func (s *Server) WishlistsEditHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return views.Render(c, views.WishlistEditView(domain.ToWishlist(wishlist)))
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
		ID:          wishlist.ID,
		Name:        name,
		Description: description,
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

	wishlistsView := utils.Map(wishlists, func(wishlist gateway.Wishlist) domain.Wishlist {
		return domain.ToWishlist(wishlist)
	})

	return views.Render(c, views.WishlistIndexView(
		domain.WishlistIndex{
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
		ID:          id,
		Name:        name,
		Description: description,
		ShareCode:   shareId,
		Public:      false,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating wishlist: %s", err))
	}

	return c.Redirect(http.StatusFound, "/admin/wishlists")
}

func (s *Server) ItemsNewHandler(c echo.Context) error {
	id := c.Param("id")

	wishlist, err := s.queries.FindWishlist(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return views.Render(c, views.ItemNewView(domain.ToWishlist(wishlist)))
}

func (s *Server) ItemsPostHandler(c echo.Context) error {
	wishlistId := c.FormValue("wishlist_id")

	wishlist, err := s.queries.FindWishlist(s.ctx, wishlistId)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	name := c.FormValue("name")
	link := c.FormValue("link")
	description := c.FormValue("description")
	imageURL := c.FormValue("image_url")
	quantity, err := strconv.ParseInt(c.FormValue("quantity"), 10, 64)
	floatPrice, err := strconv.ParseFloat(c.FormValue("price"), 64)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error parsing quantity or price: %s", err))
	}

	moneyPrice := money.NewFromFloat(floatPrice, money.USD)
	intPrice := moneyPrice.Amount()

	id, err := GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error generating id: %s", err))
	}

	err = s.queries.CreateWishlistItem(s.ctx, gateway.CreateWishlistItemParams{
		ID:          id,
		WishlistID:  wishlistId,
		Name:        name,
		Description: description,
		Link:        link,
		ImageUrl:    imageURL,
		Quantity:    quantity,
		Price:       intPrice,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating wishlist: %s", err))
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/admin/wishlists/%s", wishlist.ID))
}

func (s *Server) ItemsEditHandler(c echo.Context) error {
	id := c.Param("id")

	log.Info("id: ", id)

	item, err := s.queries.FindItem(s.ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting item: %s", err))
	}

	wishlist, err := s.queries.FindWishlist(s.ctx, item.WishlistID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	return views.Render(c, views.ItemEditView(domain.ToWishlist(wishlist), domain.ToItem(item)))
}
