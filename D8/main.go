package main

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"hacktiv/handler"
	"hacktiv/middleware"
)

func main() {
	e := echo.New()

	// built-in middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Gzip())

	// custom middleware
	e.Use(middleware.WithAuth)

	e.POST("/login", handler.LoginHandler)

	e.GET("/panic", func(c echo.Context) error {
		panic("panic")
	})

	RegisterUserRoutes(e)
	RegisterProductRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func RegisterUserRoutes(e *echo.Echo) {
	u := e.Group("/users")
	u.GET("", handler.UserHandler)
}

func RegisterProductRoutes(e *echo.Echo) {
	p := e.Group("/products")
	p.POST("", handler.ProductCreateHandler)
	p.GET("/:id", handler.ProductHandler)
}
