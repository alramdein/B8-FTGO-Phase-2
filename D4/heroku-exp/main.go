package main

import (
	"net/http" // MUX
	"os"

	_ "github.com/alramdein/heroku-example/docs" // Import generated docs

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// @title API FTGO Batch 8
// @version 1.0
// @description Really high-performance API FTGO Batch 8
// @host localhost:8080
// @BasePath /
func main() {
	mux := http.NewServeMux()
	e := echo.New()

	e.GET("/users", UserHandler)

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8112" // default
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

// List Users lists all existing users
//
//	@Summary      List users
//	@Description  get users
//	@Tags         users
//	@Accept       json
//	@Produce      json
//	@Param        q    query     string  false  "name search by q"  Format(email)
//	@Success      200  {array}   User
//	@Failure      400  {object}  HTTPError
//	@Failure      404  {object}  HTTPError
//	@Failure      500  {object}  HTTPError
//	@Router       /users [get]
func UserHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": []interface{}{
			map[string]interface{}{
				"name":         "Alif",
				"email":        "alif@go.dev",
				"ini_dari_cfg": os.Getenv("INI_DARI_CFG"),
			},
		},
	})
}
