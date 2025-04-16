package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// Login handler
func Login(c echo.Context) error {
	// Get username and password from the request
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	log.Println("Login request:", req)

	// Validate credentials (this is just a simple example, in a real application you would check against a database)
	if req.Username != "admin" || req.Password != "admin" {
		return echo.ErrUnauthorized
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return echo.ErrInternalServerError
	}

	// Return the token
	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
