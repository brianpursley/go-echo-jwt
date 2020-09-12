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
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handlers.JwtKey),
		Skipper:    skipAuthentication,
	}))

	// Health check
	e.GET("/health", handlers.GetHealthCheck)

	// Token routes
	e.POST("/auth/token", handlers.PostAuthToken)
	e.POST("/auth/token/refresh", handlers.PostAuthTokenRefresh)

	// Authenticated routes
	e.GET("/hello", handlers.GetHello)

	e.Logger.Fatal(e.Start(":8000"))
}

func skipAuthentication(c echo.Context) bool {
	unauthenticatedPaths := []string{
		"/health",
		"/auth/token",
		"/auth/token/refresh",
	}
	for _, p := range unauthenticatedPaths {
		if c.Path() == p {
			return true
		}
	}
	return false
}
