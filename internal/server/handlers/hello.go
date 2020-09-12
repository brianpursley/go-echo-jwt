package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

func GetHello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	department := claims["department"].(string)
	role := claims["role"].(string)
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s. Your department is %q and your role is %q.", username, department, role))
}
