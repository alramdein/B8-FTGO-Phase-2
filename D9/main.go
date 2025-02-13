package main

import (
	"fmt"
	"hacktiv/handler"
	"hacktiv/model"
	"hacktiv/repository"
	"hacktiv/usecase"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(ComposeConnStr()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRole{})

	e := echo.New()

	// built-in middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Gzip())

	dbTransactioner := repository.NewDBTransactioner(db)
	userRepo := repository.NewUserRepository(db)
	userUseacase := usecase.NewUserUsecase(userRepo, dbTransactioner)
	userHandler := handler.NewUserHandler(userUseacase)

	userHandler.RegisterUserRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func ComposeConnStr() string {
	return fmt.Sprintf(`host=%s user=%s  password=%s  dbname=%s  port=%s  sslmode=%s  TimeZone=%s`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)
}
