package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"whishlist/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/js/*", echo.WrapHandler(fileServer))

	e.GET("/_ping", s.PingHandler)

	wishlists := web.NewWishlists(s.ctx, s.queries)
	e.GET("/wishlists/:id", wishlists.WishlistShowHandler)

	return e
}

func (s *Server) PingHandler(c echo.Context) error {
	resp := map[string]string{
		"status": "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
