package main

import (
	"net/http"

	"github.com/dwmccheyne/typey.git/internal/handlers"
	"github.com/dwmccheyne/typey/internal/repository"
)

func main() {
	// This is a placeholder implementation
	// You should replace this with your actual database connection and HTTP server setup

	repo := repository.NewDnaSequenceRepository() // replace with real implementation
	handler := handlers.DnaHandler{
		Repository: repo,
	}

	http.HandleFunc("/sequences", handler.GetSequence)

	http.ListenAndServe(":8080", nil)
}
