package repository

import (
	"github.com/dwmccheyne/typey/internal/models"
)

type DnaSequenceRepository interface {
	Save(seq models.DnaSequence) (models.DnaSequence, error)
	FindByID(id string) (models.DnaSequence, error)
	// Define more methods as needed, e.g. Delete, Update, etc.
}
