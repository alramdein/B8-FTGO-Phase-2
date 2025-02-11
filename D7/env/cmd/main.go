package main

import (
	"fmt"
	"os"
	"reflect"

	// "github.com/joho/godotenv" // implementasi menggunakan godotenv.Load()
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // Ini autload .env file di current directory
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getenv("VALUE1"))
	fmt.Println(os.Getenv("VALUE2"))
	fmt.Println(reflect.TypeOf(os.Getenv("VALUE2")))
}
