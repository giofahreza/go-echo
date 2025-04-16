package middleware

import (
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(emw.Logger())
	e.Use(emw.Recover())
	// e.Use(AuthMiddleware)
}
