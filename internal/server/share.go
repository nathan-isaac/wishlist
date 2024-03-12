package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"wishlist/internal/domain"
	"wishlist/internal/views"
)

func (s *Server) ShareShowHandler(c echo.Context) error {
	code := c.Param("code")

	wishlist, err := s.queries.FindWishlistByShareCode(s.ctx, code)

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

	share := domain.Share{
		Id:       wishlist.ID,
		Code:     code,
		Wishlist: domain.ToWishlist(wishlist),
		Items:    wishlistItems,
	}

	return views.Render(c, views.ShareView(share))
}
