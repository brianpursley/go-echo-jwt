package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"webapp1/internal/server/handlers"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Unauthenticated routes
	e.GET("/hello", handlers.GetHello)

	// Token routes
	e.POST("/auth/token", handlers.PostAuthToken)
	e.POST("/auth/token/refresh", handlers.PostAuthTokenRefresh)

	// Authenticated routes
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte(handlers.JwtKey)))
	r.GET("/hello", handlers.GetRestrictedHello)

	e.Logger.Fatal(e.Start(":8000"))
}
