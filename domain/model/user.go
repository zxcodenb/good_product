// Package model contains the domain models
package model

import "gorm.io/gorm"

// User represents a user entity
type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"column:id;type:varchar(36);primaryKey"` // Unique identifier
	PhoneNo  string `json:"phone_no" gorm:"column:phone_no;type:varchar(20);uniqueIndex;not null"` // Phone number
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null"` // Password
	Name     string `json:"name" gorm:"column:name;type:varchar(100)"` // User name
	Remark   string `json:"remark" gorm:"column:remark;type:varchar(255)"` // Remark
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "users"
}