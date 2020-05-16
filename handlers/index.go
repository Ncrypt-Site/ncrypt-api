package handlers

import (
	"github.com/labstack/echo"
)

func (di *DI) Index(c echo.Context) error {
	return c.Redirect(301, "/api/v1")
}
