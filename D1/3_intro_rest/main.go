package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserFormData struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"users": []interface{}{
				map[string]interface{}{
					"name":  "Alif",
					"email": "alif@go.dev",
				},
			},
		})
	})

	e.POST("/users", func(c echo.Context) error {
		// var payload UserPayload // declaration
		// payload := UserPayload{} // i
		// payload := new(UserPayload)
		payload := &UserPayload{} // instantiate + point
		if err := c.Bind(payload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"users": []interface{}{
					map[string]interface{}{
						"message": "Invalid payload",
					},
				},
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"users": []interface{}{
				map[string]interface{}{
					"message": "User created",
					"user":    payload,
				},
			},
		})
	})

	e.POST("/users/form", func(c echo.Context) error {
		payload := UserFormData{}
		payload.Name = c.FormValue("name")
		payload.Email = c.FormValue("email")

		return c.JSON(http.StatusOK, map[string]interface{}{
			"users": []interface{}{
				map[string]interface{}{
					"message": "User created",
					"user":    payload,
				},
			},
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
