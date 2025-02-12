package handler

import (
	"hacktiv/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// PASTIKAN JANGAN DATA CONFEDENTIAL
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return c.JSON(http.StatusInternalServerError, model.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error secret is empty",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.LoginResponse{
		Token: tokenString,
	})
}
