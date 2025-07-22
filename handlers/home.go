package handlers

import (
	"net/http"

	"www.example.com/met/components"
)

// Home renders the main landing page
func Home(w http.ResponseWriter, r *http.Request) {
	// Create the home component
	component := components.Home()

	// Render the component
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering home page", http.StatusInternalServerError)
		return
	}
}
