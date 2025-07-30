package api

import (
	"item-value/common/response"
	"item-value/constants"
	"item-value/domain/dto"
	"item-value/domain/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	response.BaseResponse
}

// CreateUser
func (ctrl *UserController) CreateUser(c *gin.Context) {

	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Error(c, constants.ErrCodeBadRequest, "请求参数错误:"+err.Error())
		return
	}

	userService := service.NewUserService()
	user, err := userService.CreateUser(req)
	if err != nil {
		ctrl.Error(c, constants.ErrCodeBadRequest, "创建用户失败:"+err.Error())
		return
	}

	ctrl.Success(c, user)
}

// Login
func (ctrl *UserController) Login(c *gin.Context) {

	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.Error(c, constants.ErrCodeBadRequest, "请求参数错误:"+err.Error())
		return
	}

	service := service.NewUserService()
	response, err := service.Login(req)
	if err != nil {
		ctrl.Error(c, constants.ErrCodeInternal, err.Error())
		return
	}

	ctrl.Success(c, response)
}

// TODO: update user
// func (ctrl *UserController) UpdateUser(c *gin.Context) {
// 	var req dto.UserUpdateRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		ctrl.Error(c, constants.ErrCodeBadRequest, "请求参数错误:"+err.Error())
// 		return
// 	}
// 	response, err := service.UpdateUser(req)
// }
