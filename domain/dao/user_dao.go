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
// 现在是 UserDAO 的一个方法
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

// CreateUser 在数据库中创建一个新用户
// 现在是 UserDAO 的一个方法
func (ud *UserDAO) CreateUser(user *model.User) error {
	// 使用结构体内的 db 实例
	result := ud.db.Create(user)
	return result.Error
}

func (ud *UserDAO) CheckUserExists(phoneNo string, name string) (bool, error) {

	var count int64
	//根据手机号和姓名查询用户
	result := ud.db.Where("phone_no = ? AND name = ?", phoneNo, name).First(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
