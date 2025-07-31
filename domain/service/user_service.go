package service

import (
	"errors"
	"item-value/domain/dao"
	"item-value/domain/dto"
	"item-value/domain/model"
	"item-value/utils"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		userDAO: dao.NewUserDAO(db),
	}
}

// Login 登录
func (s *UserService) Login(req dto.UserLoginRequest) (*dto.LoginResponse, error) {

	// 校验用户名
	user, err := s.userDAO.GetUserByphone(req.PhoneNo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("手机号或密码错误")
		}
		return nil, err
	}

	if !user.VerifyPassword(req.Password) {
		return nil, errors.New("手机号或密码错误")
	}

	//todo: 生成token
	// 生成Token（这里简化处理，实际应该使用JWT）
	token := "user_" + user.ID + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	expiresAt := time.Now().Add(24 * time.Hour) // 24小时过期

	return &dto.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// CreateUser 创建用户
func (s *UserService) CreateUser(req dto.CreateUserRequest) (*dto.UserResponse, error) {

	// 检查用户是否存在
	hasUser, _ := s.userDAO.CheckUserExists(req.Username)
	if hasUser {
		return nil, errors.New("该手机号绑定用户已存在")
	}

	// 创建用户
	user := model.User{
		ID:         utils.GenerateSnowflakeID(),
		Name:       req.Username,
		Phone:      req.Phone,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Remark:     req.Remark,
	}
	// 密码加密
	if err := user.SetPassword(req.Password); err != nil {
		return nil, errors.New("错误，请联系管理员")
	}
	// 保存用户
	err := s.userDAO.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	userRes := dto.UserResponse{

		PhoneNo:    user.Phone,
		Name:       user.Name,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}

	return &userRes, nil

}
