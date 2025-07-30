// Package controller handles HTTP requests
package controller

import (
	"encoding/json"
	"item-value/domain/dto"
	"item-value/domain/service"
	"net/http"

	"github.com/gorilla/mux"
)

// ItemController handles item-related HTTP requests
type ItemController struct {
	itemService *service.ItemService
}

// NewItemController creates a new ItemController instance
func NewItemController() *ItemController {
	return &ItemController{
		itemService: service.NewItemService(),
	}
}

// RegisterRoutes registers item-related routes
func (c *ItemController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/items", c.CreateItem).Methods("POST")
	router.HandleFunc("/items", c.ListItems).Methods("GET")
	router.HandleFunc("/items/{id}", c.GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", c.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", c.DeleteItem).Methods("DELETE")
}

// CreateItem handles the creation of a new item
func (c *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
	var req dto.ItemCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Name == "" {
		http.Error(w, "Item name is required", http.StatusBadRequest)
		return
	}

	response, err := c.itemService.CreateItem(req)
	if err != nil {
		http.Error(w, "Failed to create item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetItem handles retrieving an item by its ID
func (c *ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := c.itemService.GetItem(id)
	if err != nil {
		http.Error(w, "Item not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateItem handles updating an existing item
func (c *ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req dto.ItemUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	response, err := c.itemService.UpdateItem(id, req)
	if err != nil {
		http.Error(w, "Failed to update item: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteItem handles deleting an item by its ID
func (c *ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.itemService.DeleteItem(id); err != nil {
		http.Error(w, "Failed to delete item: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListItems handles retrieving all items
func (c *ItemController) ListItems(w http.ResponseWriter, r *http.Request) {
	response, err := c.itemService.ListItems()
	if err != nil {
		http.Error(w, "Failed to retrieve items: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}