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

	e.GET("/checkout/:id", s.CheckoutShowHandler)
	e.POST("/checkout", s.CheckoutCreateHandler)
	e.PUT("/checkout/:id", s.CheckoutUpdateHandler)

	admin := e.Group("/admin")

	admin.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username string, password string, context echo.Context) (bool, error) {
			return username == s.admin.Username && password == s.admin.Password, nil
		},
	}))

	admin.GET("/lists", s.ListsIndexHandler)
	admin.POST("/lists", s.ListsCreateHandler)
	admin.GET("/lists/new", s.ListsNewHandler)
	admin.GET("/lists/:id", s.ListsShowHandler)
	admin.GET("/lists/:id/edit", s.ListsEditHandler)
	admin.PUT("/lists/:id", s.ListsUpdateHandler)
	admin.DELETE("/lists/:id", s.ListsDeleteHandler)

	admin.GET("/lists/:id/items/new", s.ItemsNewHandler)
	admin.POST("/items", s.ItemsCreateHandler)
	admin.GET("/items/:id/edit", s.ItemsEditHandler)
	admin.PUT("/items/:id", s.ItemsUpdateHandler)
	admin.DELETE("/items/:id", s.ItemsDeleteHandler)

	return e
}

func (s *Server) PingHandler(c echo.Context) error {
	resp := map[string]string{
		"status": "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
