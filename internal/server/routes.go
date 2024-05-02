package server

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"wishlist/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if s.sentryEnabled {
		e.Use(sentryecho.New(sentryecho.Options{
			Repanic: true,
		}))
	}

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/_ping", s.PingHandler)

	e.GET("/shares/:share_code", s.SharesShowHandler)

	e.GET("/checkouts/:checkout_id", s.CheckoutsShowHandler)
	e.POST("/checkouts", s.CheckoutsCreateHandler)
	e.PUT("/checkouts/:checkout_id", s.CheckoutsUpdateHandler)
	e.PUT("/checkout-items/:checkout_item_id", s.CheckoutItemsUpdateHandler)

	admin := e.Group("/admin")

	admin.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username string, password string, context echo.Context) (bool, error) {
			return username == s.admin.Username && password == s.admin.Password, nil
		},
	}))

	admin.GET("/lists", s.ListsIndexHandler)
	admin.POST("/lists", s.ListsCreateHandler)
	admin.GET("/lists/new", s.ListsNewHandler)
	admin.GET("/lists/:list_id", s.ListsShowHandler)
	admin.GET("/lists/:list_id/edit", s.ListsEditHandler)
	admin.PUT("/lists/:list_id", s.ListsUpdateHandler)
	admin.DELETE("/lists/:list_id", s.ListsDeleteHandler)

	admin.GET("/lists/:list_id/items/new", s.ItemsNewHandler)
	admin.POST("/items", s.ItemsCreateHandler)
	admin.GET("/items/:list_item_id/edit", s.ItemsEditHandler)
	admin.PUT("/items/:list_item_id", s.ItemsUpdateHandler)
	admin.DELETE("/items/:list_item_id", s.ItemsDeleteHandler)

	return e
}

func (s *Server) PingHandler(c echo.Context) error {
	resp := map[string]string{
		"status": "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
