package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func (di *DI) Index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/api/v1")
}
