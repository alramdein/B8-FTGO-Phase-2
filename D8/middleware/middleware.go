package middleware

import (
	"fmt"
	"hacktiv/model"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Whitelist di middleware auth
		if c.Request().URL.Path == "/login" {
			return next(c)
		}

		header := c.Request().Header
		fmt.Println(header)
		auth := strings.TrimPrefix(header.Get("Authorization"), "Bearer ")
		fmt.Println(auth)
		if auth == "" {
			fmt.Println("Error auth is empty")
			return c.JSON(http.StatusUnauthorized, model.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			fmt.Println("Error secret is empty")
			return c.JSON(http.StatusUnauthorized, model.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}

		token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			fmt.Println("Error failed to parse token")
			return c.JSON(http.StatusUnauthorized, model.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Error failed to map token")
			return c.JSON(http.StatusUnauthorized, model.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}

		// Paling penting
		if !token.Valid {
			fmt.Println("Error token invalid")
			return c.JSON(http.StatusUnauthorized, model.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}

		fmt.Println("CLAIMS: ", claims)

		// Pengeeckkan Authorization (hak akses)
		// ....

		c.Set("username", claims["username"])
		// masukin informasi login lain asal jangan confidential

		return next(c)
	}
}
