package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

type User3 struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   *int32 `json:"age"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func main() {
	router := httprouter.New()
	// Handlers
	// Contoh response raw text
	router.GET("/users", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "List of users")
	})

	// Contoh JSON
	router.POST("/users/json/bad", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name": "Alif", "email": "alif@go.dev"}`)
	})

	router.PATCH("/users/json/good/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		v := p.ByName("userNameForUnregisteredNumber")

		user := User3{
			Name:  v,
			Email: v + "@go.dev",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	router.POST("/users", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var user User3

		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	router.POST("/panic", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("database error.")
	})

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusInternalServerError)

		log := logrus.New()
		log.SetFormatter(&ecslogrus.Formatter{})

		log2 := log.WithFields(logrus.Fields{
			"id":         "234234",
			"handler":    "panic handler",
			"metadata-1": "value-1",
			"metadata-2": "value-2",
			"line":       "fungsiA:10",
		})

		log2.Error(i)

		user := ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "something went wrong", // obfuscate error

		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}

	fmt.Println("Server running on port 3000...")
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println(server.ListenAndServe().Error())
}
