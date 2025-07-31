// Package dto contains data transfer objects
package dto

import "time"

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" example:"newuser"`
	Password string `json:"password" binding:"required,min=6,max=50" example:"password123"`
	Phone    string `json:"phone" binding:"omitempty,max=20" example:"13800138000"`
	Remark   string `json:"remark"` // Remark
}

// UserUpdateRequest represents the request body for updating a user
type UserUpdateRequest struct {
	Password string `json:"password"` // Password
	Name     string `json:"name"`     // User name
	Remark   string `json:"remark"`   // Remark
}

// UserResponse represents the response body for a user
type UserResponse struct {
	PhoneNo    string    `json:"phone"`        // Phone number
	Name       string    `json:"name"`         // User name
	Remark     string    `json:"remark"`       // Remark
	CreateTime time.Time `json:"created_time"` // Creation time
	UpdateTime time.Time `json:"updated_time"` // Update time
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// user Login request
type UserLoginRequest struct {
	PhoneNo  string `json:"phone" binding:"required"`    // Phone number
	Password string `json:"password" binding:"required"` // Password
}

// UserListResponse represents the response body for a list of users
type UserListResponse struct {
	Users []UserResponse `json:"users"` // List of users
}
