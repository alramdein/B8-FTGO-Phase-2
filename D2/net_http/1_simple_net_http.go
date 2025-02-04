package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main1() {
	server := http.Server{
		Addr: "localhost:3000",
	}

	// Handlers
	// Contoh response raw text
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "List of users")
	})

	// Contoh JSON
	// -- DONT DO THIS
	http.HandleFunc("/users/json/bad", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name": "Alif", "email": "alif@go.dev"}`)
	})

	// Better solution!
	http.HandleFunc("/users/json/good", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Name:  "Yosua",
			Email: "yosua@go.dev",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	fmt.Println("Server running on port 3000...")
	err := server.ListenAndServe() // blocking
	if err != nil {
		panic(err)
	}
}
