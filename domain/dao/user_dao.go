package dao

import (
	"item-value/domain/model"
	"item-value/utils"

	"gorm.io/gorm"
)

// UserDAO 结构体用于封装与用户相关的数据库操作
// 它包含一个数据库连接实例
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 是一个构造函数，用于创建一个新的 UserDAO 实例
// 它接收一个 *gorm.DB 连接作为参数
func NewUserDAO() *UserDAO {

	return &UserDAO{
		db: utils.GetDB(),
	}

}

// GetUserByID 根据ID从数据库中查找用户
func (ud *UserDAO) GetUserByID(id string) (*model.User, error) {
	var user model.User
	// 使用结构体内的 db 实例，而不是全局的 config.DB
	result := ud.db.Where("id = ?", id).First(&user)

	// result.Error 会在发生错误时（包括未找到记录）返回一个 error
	if result.Error != nil {
		return nil, result.Error
	}

	// 如果没有错误，则返回找到的用户和 nil
	return &user, nil
}

func (ud *UserDAO) CreateUser(user *model.User) error {
	// 使用结构体内的 db 实例
	result := ud.db.Create(user)
	return result.Error
}

// CheckUserExists 检查用户是否存在
func (ud *UserDAO) CheckUserExists(phoneNo string) (bool, error) {

	var count int64
	//根据手机号查询用户
	result := ud.db.Where("phone = ?", phoneNo).First(&count)
	err := result.Count(&count).Error
	return count > 0, err

}

// GetUserByPhone 根据phone获取用户
func (ud *UserDAO) GetUserByphone(phone string) (*model.User, error) {

	var user model.User
	//根据name查询用户
	result := ud.db.Where("name = ?", phone).First(&user)
	return &user, result.Error

}
