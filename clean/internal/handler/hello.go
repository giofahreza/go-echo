package handler

import (
	"net/http"

	"go-echo/internal/model"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, model.HelloResponse{
		Message: "Hello World!",
	})
}
