package server

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"wishlist/internal/views/checkout"
)

func (s *Server) CheckoutShowHandler(c echo.Context) error {
	id := c.Param("id")

	slog.Info("checkout show handler", slog.String("id", id))

	return Render(c, checkout.CheckoutShowView())
}

func (s *Server) CheckoutCreateHandler(c echo.Context) error {
	wishlistId := c.FormValue("wishlist_id")
	itemId := c.FormValue("wishlist_item_id")

	slog.Info("checkout show handler", slog.String("wishlist_id", wishlistId), slog.String("wishlist_item_id", itemId))

	// create checkout database records
	// redirect to show view

	return Render(c, checkout.CheckoutShowView())
}

// how does a user cancel a checkout?
// maybe we need a confirmation page?
