package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"

	_ "github.com/joho/godotenv/autoload"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Middleware1(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Do something

		// Auhtentication

		// Logging

		// Error handler dll

		fmt.Println("Middleware 1 called..")
		r.Header.Add("X-Intercepted-By", "Middleware1")
		next(w, r, p)
	}
}

func Middleware2(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Do something
		fmt.Println("Middleware 2 called..")
		r.Header.Add("X-Intercepted-By-2", "Middleware2")
		next(w, r, p)
	}
}

func MovieHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Intercepted req header: ")
	fmt.Printf("%v", r.Header)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Success"))
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "errornya pake json yang rapih", http.StatusBadRequest)
		return
	}

	// PASTIKAN JANGAN DATA CONFEDENTIAL
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		http.Error(w, "Error secret is empty", http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"token": tokenString,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func WithAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Do something
		fmt.Println("WithAuth called..")
		fmt.Println(r.Header)
		auth := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		fmt.Println(auth)
		if auth == "" {
			fmt.Println("Error auth is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			fmt.Println("Error secret is empty")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			fmt.Println("Error failed to parse token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Error failed to map token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Paling penting
		if !token.Valid {
			fmt.Println("Error token invalid")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println("CLAIMS: ", claims)

		// Pengeeckkan Authorization (hak akses)
		// ....

		next(w, r, p)
	}
}

func Protected(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Hal2 yang berhububgan dangen endpoint nya
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Success"))
}
func Handler1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Hal2 yang berhububgan dangen endpoint nya
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Success"))
}
func Handler2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Hal2 yang berhububgan dangen endpoint nya
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Success"))
}

func main() {
	router := httprouter.New()

	router.GET("/movies", Middleware2(Middleware1(MovieHandler)))
	router.GET("/movies/:id", MovieHandler)

	router.POST("/login", Login)

	router.GET("/protected", WithAuth(Protected))
	router.GET("/protected", WithAuth(Handler1))
	router.GET("/protected", WithAuth(Handler2))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server running on port 3000...")
	server.ListenAndServe()
}
