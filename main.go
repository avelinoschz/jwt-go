package main

import (
	"log"
	"net/http"

	"github.com/avelinoschz/jwt-go/authentication"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", authentication.Login)
	mux.HandleFunc("/validate", authentication.ValidateToken)

	log.Println("Server listening request on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
