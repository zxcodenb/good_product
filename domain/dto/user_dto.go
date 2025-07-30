// Package dto contains data transfer objects
package dto

// UserCreateRequest represents the request body for creating a user
type UserCreateRequest struct {
	PhoneNo  string `json:"phone_no" binding:"required"` // Phone number
	Password string `json:"password" binding:"required"` // Password
	Name     string `json:"name"`                        // User name
	Remark   string `json:"remark"`                      // Remark
}

// UserUpdateRequest represents the request body for updating a user
type UserUpdateRequest struct {
	Password string `json:"password"` // Password
	Name     string `json:"name"`     // User name
	Remark   string `json:"remark"`   // Remark
}

// UserResponse represents the response body for a user
type UserResponse struct {
	ID        string `json:"id"`         // Unique identifier
	PhoneNo   string `json:"phone_no"`   // Phone number
	Name      string `json:"name"`       // User name
	Remark    string `json:"remark"`     // Remark
	CreatedAt string `json:"created_at"` // Creation time
	UpdatedAt string `json:"updated_at"` // Update time
}

// UserListResponse represents the response body for a list of users
type UserListResponse struct {
	Users []UserResponse `json:"users"` // List of users
}
