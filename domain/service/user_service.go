package service

import (
	"item-value/domain/dao"
	"item-value/domain/dto"
	"item-value/domain/model"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		userDAO: dao.NewUserDAO(),
	}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(req dto.UserCreateRequest) (*dto.UserResponse, error) {

	// 检查用户是否存在
	count, err := s.userDAO.CheckUserExists(req.PhoneNo, req.Name)

	user := model.User{
		Name:     req.Name,
		Password: req.Password,
		PhoneNo:  req.PhoneNo,
		Remark:   req.Remark,
	}
	err := s.userDAO.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return nil
}
