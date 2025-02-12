package handler

import (
	"fmt"
	"net/http"

	"hacktiv/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func UserHandler(c echo.Context) error {
	fmt.Println(c.Request().Header)
	name := c.QueryParam("name")

	fmt.Println("INI USERNAME USER: ", c.Get("username"))

	return c.JSON(http.StatusOK, model.UserResponse{
		Name:  name,
		Email: "alif@go.dev",
	})
}

func ProductHandler(c echo.Context) error {
	id := c.Param("id")

	// masuk usecase
	// r.usecase.GetProduct(id)

	return c.JSON(http.StatusOK, model.Product{
		ID:   id,
		Name: "Product GET",
	})
}

func ProductCreateHandler(c echo.Context) error {
	var product model.Product
	err := c.Bind(&product)
	if err != nil {
		fmt.Println(err.Error()) // selalu log errornya
		return c.JSON(http.StatusBadRequest, model.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// ini terjadi di usecase
	product.ID = uuid.NewString()

	return c.JSON(http.StatusOK, product)
}
