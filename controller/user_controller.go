// Package controller handles HTTP requests
package controller

import (
	"encoding/json"
	"item-value/domain/dto"
	"item-value/domain/service"
	"net/http"

	"github.com/gorilla/mux"
)

// UserController handles user-related HTTP requests
type UserController struct {
	userService *service.UserService
}

// NewUserController creates a new UserController instance
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// RegisterRoutes registers user-related routes
func (c *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", c.CreateUser).Methods("POST")
	router.HandleFunc("/users", c.ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", c.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", c.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", c.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/phone/{phoneNo}", c.GetUserByPhoneNo).Methods("GET")
}

// CreateUser handles the creation of a new user
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.PhoneNo == "" {
		http.Error(w, "Phone number is required", http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	response, err := c.userService.CreateUser(req)
	if err != nil {
		if _, ok := err.(*service.UserAlreadyExistsError); ok {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetUser handles retrieving a user by its ID
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := c.userService.GetUser(id)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUserByPhoneNo handles retrieving a user by phone number
func (c *UserController) GetUserByPhoneNo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phoneNo := vars["phoneNo"]

	response, err := c.userService.GetUserByPhoneNo(phoneNo)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateUser handles updating an existing user
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req dto.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	response, err := c.userService.UpdateUser(id, req)
	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteUser handles deleting a user by its ID
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.userService.DeleteUser(id); err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListUsers handles retrieving all users
func (c *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	response, err := c.userService.ListUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}