package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
