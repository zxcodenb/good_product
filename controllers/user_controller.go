package controllers

import (
	"item-value/common/response"
	"item-value/constants"
	"item-value/domain/dto"
	"item-value/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserAPI 结构体，负责处理用户相关的 API 请求
type UserController struct {
	response.BaseResponse
	userService *service.UserService
}

// NewUserAPI 创建一个新的 UserAPI 实例
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		userService: service.NewUserService(db),
	}
}

// CreateUser
func (api *UserController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, constants.ErrCodeBadRequest, "请求参数错误:"+err.Error())
		return
	}

	user, err := api.userService.CreateUser(req)
	if err != nil {
		api.Error(c, constants.ErrCodeBadRequest, "创建用户失败:"+err.Error())
		return
	}

	api.Success(c, user)
}

// Login 处理用户登录的请求
func (api *UserController) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, constants.ErrCodeBadRequest, "请求参数错误:"+err.Error())
		return
	}

	response, err := api.userService.Login(req)
	if err != nil {
		api.Error(c, constants.ErrCodeInternal, err.Error())
		return
	}

	api.Success(c, response)
}
