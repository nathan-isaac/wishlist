package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HxRedirect(c echo.Context, url string) error {
	c.Response().Header().Set("HX-Redirect", url)
	c.Response().WriteHeader(http.StatusFound)
	return nil
}
