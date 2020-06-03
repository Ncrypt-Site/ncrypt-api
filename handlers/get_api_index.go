package handlers

import (
	"github.com/labstack/echo"
	"ncrypt-api/helpers"
)

//GetIndexV1 handles GET /api/v1
func (di *DI) GetIndexV1(c echo.Context) error {
	return c.JSON(200, helpers.BuildResponse(200, "nCrypt API v: 1.0.0", nil, nil, nil))
}
