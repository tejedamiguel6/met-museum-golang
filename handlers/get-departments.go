package handlers

import (
	"net/http"

	"www.example.com/met/components"
	"www.example.com/met/services"
)

// getting department list form component

func DepartmentList(w http.ResponseWriter, r *http.Request) {
	// calling the api
	departments, err := services.CallMetAPI()
	if err != nil {
		http.Error(w, "Failed to fetch departments", http.StatusInternalServerError)
		return
	}

	// Create the DepartmentList component with fetched data
	component := components.DepartmentList(departments)

	// Render the component
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering component", http.StatusInternalServerError)
	}

}
