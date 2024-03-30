package main

import (
	// New import

	"snippetbox.alexedwards.net/internal/models"
)

// Include a Snippets field in the templateData struct.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
