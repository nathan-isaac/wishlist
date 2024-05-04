package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"wishlist/internal/domain"
	"wishlist/internal/utils"
	"wishlist/internal/views/share"
)

func findCheckoutId(c echo.Context) (string, error) {
	cookieId, err := c.Cookie(CHECKOUT_ID_COOKIE_NAME)

	if err != nil {
		return "", err
	}

	if cookieId.Value != "" {
		return cookieId.Value, nil
	}

	return c.QueryParam("checkoutId"), nil
}

func (s *Server) SharesShowHandler(c echo.Context) error {
	code := c.Param("share_code")

	wishlist, err := s.queries.FindListByShareCode(s.ctx, code)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist: %s", err))
	}

	items, err := s.queries.FilerItemsForList(s.ctx, wishlist.ListID)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting wishlist items: %s", err))
	}

	checkoutId, err := findCheckoutId(c)

	if err != nil {
		slog.Warn("error getting checkoutId: %s", err)
	}

	checkoutItems, err := s.queries.FilterCheckoutItems(s.ctx, checkoutId)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error getting checkout items: %s", err))
	}

	checkoutURL := ""

	if checkoutId != "" {
		checkoutURL = fmt.Sprintf("/checkouts/%s", checkoutId)
	}

	return Render(c, share.ShareView(domain.Share{
		Id:             wishlist.ListID,
		Code:           code,
		List:           domain.ToList(wishlist),
		Items:          utils.Map(items, domain.ToItem),
		PurchasedCount: len(checkoutItems),
		CheckoutUrl:    checkoutURL,
		CheckoutId:     checkoutId,
	}))
}
