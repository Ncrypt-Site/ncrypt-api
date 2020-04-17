package handlers

import (
	"github.com/labstack/echo"
	"ncrypt-api/helpers"
)

func (di DI) GetIndexV1(c echo.Context) error {
	return c.JSON(200, helpers.BuildResponse(200, "NCrypt API v: 1.0.0", nil, nil, nil))
}
