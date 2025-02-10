package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "hacktiv/docs" // Import generated docs

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample API using net/http with Swagger.
// @host localhost:8080
// @BasePath /api

// @Summary Get a sample response
// @Description Returns a JSON response
// @Produce json
// @Success 200 {object} map[string]string
// @Router /sample [get]
func sampleHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, Swagger with net/http!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Generate Swagger docs
	http.HandleFunc("/api/sample", sampleHandler)

	// Serve Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Start server
	fmt.Println(http.ListenAndServe(":3333", nil).Error())
}
