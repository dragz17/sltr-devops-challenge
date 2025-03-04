package handler

import (
	"challenge-boy-welcome/internal/service"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if strings.TrimSpace(name) == "" {
		name = "Anonymous"
	}

	message := service.GetWelcomeMessage(name)
	w.WriteHeader(http.StatusOK)

	// Check for error when writing the response
	if _, err := w.Write([]byte(message)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

// NotFoundHandler menangani invalid path
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid path, please go to /welcome/", http.StatusNotFound)
}
