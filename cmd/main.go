package main

import (
	"challenge-boy-welcome/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/welcome/{name}", handler.WelcomeHandler).Methods("GET")
	r.HandleFunc("/welcome/", handler.WelcomeHandler).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	log.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
