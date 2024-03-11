package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"whishlist/internal/views"
)

func (s *Server) ShareShowHandler(c echo.Context) error {
	code := c.Param("code")

	wishlist, err := s.queries.FindWishlistByShareCode(s.ctx, code)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	items, err := s.queries.ListWishlistsItemsForWishlist(s.ctx, wishlist.ID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist items: %s", err))
	}

	wishlistItems := make([]views.Item, len(items))

	for i, item := range items {
		wishlistItems[i] = views.Item{
			Id:                item.ID,
			Name:              item.Name,
			Link:              item.Link,
			ImageUrl:          item.ImageUrl,
			Description:       item.Description,
			Price:             fmt.Sprintf("%d", item.Price),
			NeededQuantity:    fmt.Sprintf("%d", item.Quantity),
			PurchasedQuantity: "0",
		}
	}

	share := views.Share{
		Id:       wishlist.ID,
		Code:     code,
		Wishlist: views.ToWishlist(wishlist),
		Items:    wishlistItems,
	}

	return views.Render(c, views.ShareView(share))
}
