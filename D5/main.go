package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/alramdein/template-go/handler"
	"github.com/alramdein/template-go/repository"
	"github.com/alramdein/template-go/usecase"
)

// Uncle Bob: Go Clean Architecure

func main() {
	godotenv.Load()

	// Initate datebase
	db, err := sql.Open("mysql", ComposeConnStr())
	if err != nil {
		panic(err) // boleh dipanggil panic karena connect ke db
	}

	// Depency Injection
	productRepo := repository.NewProductRepository(db)

	productUsecase := usecase.NewProductUsecase(productRepo)

	productHandler := handler.NewProductHandler(productUsecase)

	mux := http.NewServeMux()

	mux.HandleFunc("/products", productHandler.CreateProduct)

	fmt.Println("Server running on port 3000...")
	http.ListenAndServe("localhost:3000", mux) // HOST dan PORT jangan lupa dibuat ENV
}

func ComposeConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
