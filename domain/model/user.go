// Package model contains the domain models
package model

import (
	"golang.org/x/crypto/bcrypt"
)

// User represents a user entity
type User struct {
	ID         string `json:"id" gorm:"column:id;type:varchar(36);primaryKey"`                 // Unique identifier
	Phone      string `json:"phone" gorm:"column:phone;type:varchar(20);uniqueIndex;not null"` // Phone number
	Password   string `json:"password" gorm:"column:password;type:varchar(255);not null"`      // Password
	Name       string `json:"name" gorm:"column:name;type:varchar(100)"`                       // User name
	CreateTime string `json:"create_time" gorm:"column:create_time;type:datetime"`
	UpdateTime string `json:"update_time" gorm:"column:update_time;type:datetime"`
	Remark     string `json:"remark" gorm:"column:remark;type:varchar(255)"` // Remark
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "users"
}

// 密码加密
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// 密码验证
func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
