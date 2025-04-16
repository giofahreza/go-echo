package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Middleware1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 1")
		return next(c)
	}
}
func Middleware2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 2")
		return next(c)
	}
}
func Middleware3(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 3")
		return next(c)
	}
}
func Middleware4(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 4")
		return next(c)
	}
}
func Middleware5(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 5")
		return next(c)
	}
}
