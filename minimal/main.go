package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the token from the request header
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		log.Println("Token:", tokenString)

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		c.Set("user", token)
		return next(c)
	}
}

func middleware1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 1")
		return next(c)
	}
}
func middleware2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 2")
		return next(c)
	}
}
func middleware3(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 3")
		return next(c)
	}
}
func middleware4(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 4")
		return next(c)
	}
}
func middleware5(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Middleware 5")
		return next(c)
	}
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello, middleware1, middleware2, middleware3, middleware4, middleware5)
	e.POST("/login", login)

	// Protected route
	authGroup := e.Group("/auth")
	authGroup.Use(authMiddleware)

	authGroup.GET("/protected1", protected1)
	authGroup.GET("/protected2", protected2)
	authGroup.GET("/protected3/:name", protected3)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Login handler
func login(c echo.Context) error {
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

func protected1(c echo.Context) error {
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

func protected2(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ID is required")
	}
	return c.String(http.StatusOK, "Protected 2: "+id)
}

func protected3(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}
	return c.String(http.StatusOK, "Protected 3: "+name)
}
