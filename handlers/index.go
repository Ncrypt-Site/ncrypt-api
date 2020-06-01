package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

//Index handles app entry point and redirect to /api/v1
func (di *DI) Index(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/api/v1")
}
