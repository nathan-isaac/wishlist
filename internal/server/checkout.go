package server

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"time"
	"wishlist/internal/domain"
	"wishlist/internal/gateway"
	"wishlist/internal/views/checkout"
)

func (s *Server) CheckoutShowHandler(c echo.Context) error {
	id := c.Param("id")

	slog.Info("checkout show handler", slog.String("id", id))

	return Render(c, checkout.CheckoutShowView())
}

type CheckoutCreateRequest struct {
	ListId string `form:"list_id"`
	ItemId string `form:"item_id"`
}

func (s *Server) CheckoutCreateHandler(c echo.Context) error {
	var req CheckoutCreateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	slog.Info("checkout show handler", "req", req)

	list, err := s.domain.FindList(req.ListId)

	if err != nil {
		return c.String(http.StatusBadRequest, "list not found")
	}

	for _, item := range list.Items {
		if item.Id != req.ItemId {
			return c.String(http.StatusBadRequest, "item not found")
		}
	}

	checkoutId, err := domain.GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, "error generating checkout id")
	}

	err = s.queries.CreateCheckout(s.ctx, gateway.CreateCheckoutParams{
		ID:        checkoutId,
		ListID:    req.ListId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "error creating checkout")
	}

	itemId, err := domain.GenerateId()

	if err != nil {
		return c.String(http.StatusInternalServerError, "error generating checkout item id")
	}

	err = s.queries.CreateCheckoutItem(s.ctx, gateway.CreateCheckoutItemParams{
		ID:         itemId,
		CheckoutID: checkoutId,
		ListItemID: req.ItemId,
		Quantity:   1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	return HxRedirect(c, "/checkout/"+checkoutId)
}

// how does a user cancel a checkout?
// maybe we need a confirmation page?
