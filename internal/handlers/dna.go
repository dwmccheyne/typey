package handlers

import (
	"net/http"

	"github.com/dwmccheyne/typey/internal/repository"
)

type DnaHandler struct {
	Repository repository.DnaSequenceRepository
}

func (h *DnaHandler) GetSequence(w http.ResponseWriter, r *http.Request) {
	// Implement your handler here
	// Use h.Repository to interact with the database
}

// Add more methods as needed, e.g. PostSequence, DeleteSequence, etc.
