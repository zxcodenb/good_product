package controllers

import (
	"item-value/common/response"
	"item-value/constants"
	"item-value/domain/dto"
	"item-value/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct {
	response.BaseResponse
	itemService *service.ItemService
}

// NewUserController 创建一个新的 UserController 实例
func NewUItemController(db *gorm.DB) *ItemController {
	return &ItemController{
		itemService: service.NewItemService(db),
	}
}

// 创建物品
func (api *ItemController) CreateItem(c *gin.Context) {
	var req dto.ItemCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, constants.ErrCodeBadRequest, "请求参数错误:"+err.Error())
		return
	}

	isSuccess, err := api.itemService.CreateItem(req)
	if err != nil {
		api.Error(c, constants.ErrCodeBadRequest, "创建物品失败:"+err.Error())
		return
	}
	api.Success(c, isSuccess)
}

// 根据id查询物品
func (api *ItemController) GetItemById(c *gin.Context) {
	id := c.Param("id")
	item, err := api.itemService.GetItem(id)
	if err != nil {
		api.Error(c, constants.ErrCodeBadRequest, "查询物品失败:"+err.Error())
		return
	}
	api.Success(c, item)
}

// 分页查询
func (api *ItemController) List(c *gin.Context) {
	var req dto.ItemListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, constants.ErrCodeBadRequest, "请求参数错误:"+err.Error())
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	items, err := api.itemService.ListItems(&req)
	if err != nil {
		api.Error(c, constants.ErrCodeInternal, "查询物品列表失败:"+err.Error())
		return
	}

	api.Success(c, items)
}
