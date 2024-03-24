package server

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"wishlist/internal/views"
)

func (s *Server) CheckoutShowHandler(c echo.Context) error {
	id := c.Param("id")

	slog.Info("checkout show handler", slog.String("id", id))

	return views.Render(c, views.CheckoutShowView())
}
