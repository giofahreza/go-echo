package routes

import (
	"go-echo/internal/handler"
	"go-echo/internal/middleware"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	// Routes
	e.GET("/", handler.Hello, middleware.Middleware1, middleware.Middleware2, middleware.Middleware3, middleware.Middleware4, middleware.Middleware5)
	// e.GET("/", handler.Hello)
	e.POST("/login", handler.Login)

	// Protected route
	authGroup := e.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware)

	authGroup.GET("/protected1", handler.Protected1)
	authGroup.GET("/protected2", handler.Protected2)
	authGroup.GET("/protected3/:name", handler.Protected3)
}
