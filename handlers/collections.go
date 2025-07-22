package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"www.example.com/met/components"
	"www.example.com/met/services"
)

func CollectionObjectList(w http.ResponseWriter, r *http.Request) {
	collections, err := services.GetCollectionObjects()

	if err != nil {
		http.Error(w, "Failed to fetch collection objects", http.StatusInternalServerError)
		return
	}

	items, err := services.GetCollectionObjectItems(collections.ObjectIDs)
	if err != nil {
		http.Error(w, "Failed to fetch collection object details", http.StatusInternalServerError)
		return
	}

	component := components.CollectionList(items)
	// Render the component
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "Error rendering component", http.StatusInternalServerError)
	}

	fmt.Println("Collection Objects:", collections)

}

// func get the single ID
func GetCollectionObjIDItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := services.GetCollectionObjectItems([]int{id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Collection Object ID:", item)
	c.JSON(http.StatusOK, item)

}
