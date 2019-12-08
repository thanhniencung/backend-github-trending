package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func Welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to my mapp")
}
