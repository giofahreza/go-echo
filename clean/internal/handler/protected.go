package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Protected1(c echo.Context) error {
	type protected1Response struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return c.JSON(http.StatusOK, protected1Response{
		Message: "Protected 1: " + username,
		Token:   user.Raw,
	})
}

func Protected2(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
	}
	return c.String(http.StatusOK, "Protected 2: "+id)
}

func Protected3(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}
	return c.String(http.StatusOK, "Protected 3: "+name)
}
