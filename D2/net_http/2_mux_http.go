package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User2 struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main2() {
	mux := http.NewServeMux()
	// Handlers
	// Contoh response raw text
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "List of users")
	})

	// Contoh JSON
	// -- DONT DO THIS
	mux.HandleFunc("/users/json/bad", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name": "Alif", "email": "alif@go.dev"}`)
	})

	// Better solution!
	mux.HandleFunc("/users/json/good", func(w http.ResponseWriter, r *http.Request) {
		user := User2{
			Name:  "Yosua",
			Email: "yosua@go.dev",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	fmt.Println("Server running on port 3000...")
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println(server.ListenAndServe().Error())
}
