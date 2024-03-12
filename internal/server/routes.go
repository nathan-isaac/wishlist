package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"wishlist/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/_ping", s.PingHandler)

	e.GET("/share/:code", s.ShareShowHandler)

	admin := e.Group("/admin")

	admin.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username string, password string, context echo.Context) (bool, error) {
			return username == s.admin.Username && password == s.admin.Password, nil
		},
	}))

	admin.GET("/wishlists", s.WishlistsIndexHandler)
	admin.POST("/wishlists", s.WishlistsPostHandler)
	admin.GET("/wishlists/new", s.WishlistsNewHandler)
	admin.GET("/wishlists/:id", s.WishlistsShowHandler)
	admin.GET("/wishlists/:id/edit", s.WishlistsEditHandler)
	admin.PUT("/wishlists/:id", s.WishlistsUpdateHandler)
	admin.DELETE("/wishlists/:id", s.WishlistsDeleteHandler)

	admin.GET("/wishlists/:id/items/new", s.ItemsNewHandler)
	admin.POST("/items", s.ItemsPostHandler)
	admin.GET("/items/:id/edit", s.ItemsEditHandler)

	return e
}

func (s *Server) PingHandler(c echo.Context) error {
	resp := map[string]string{
		"status": "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
