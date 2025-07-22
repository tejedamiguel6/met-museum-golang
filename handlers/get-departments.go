package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func CollectionObjectList(w http.ResponseWriter, r *http.Request) {
	collectionObjects, err := services.GetCollectionObjects()
	if err != nil {
		http.Error(w, "Failed to fetch collection objects", http.StatusInternalServerError)
		return
	}

	fmt.Println("Collection Objects:", collectionObjects)

}

// func get the single ID
func GetCollectionObjIDItem(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ids, err := services.GetCollectionObjectItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Collection Object ID:", ids)
	c.JSON(http.StatusOK, ids)

}
