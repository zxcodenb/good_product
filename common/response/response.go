package response

import (
	"item-value/constants"
	"item-value/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BaseResponse 是一个包含通用响应方法的基础结构体
type BaseResponse struct {
}

// Success 成功响应
func (receiver BaseResponse) Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, dto.Response{
		Code:    constants.ErrCodeSuccess,
		Message: constants.GetErrMessage(constants.ErrCodeSuccess),
		Data:    data,
	})
}

// SuccessWithMessage 带自定义消息的成功响应
func (receiver BaseResponse) SuccessWithMessage(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, dto.Response{
		Code:    constants.ErrCodeSuccess,
		Message: message,
		Data:    data,
	})
}

// Error 错误响应
func (receiver BaseResponse) Error(c *gin.Context, code int, message string) {
	httpStatus := receiver.getHTTPStatus(code)
	if message == "" {
		message = constants.GetErrMessage(code)
	}

	c.JSON(httpStatus, dto.Response{
		Code:    code,
		Message: message,
	})
}

// ErrorWithData 带数据的错误响应
func (receiver BaseResponse) ErrorWithData(c *gin.Context, code int, message string, data any) {
	httpStatus := receiver.getHTTPStatus(code)
	if message == "" {
		message = constants.GetErrMessage(code)
	}

	c.JSON(httpStatus, dto.Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// getHTTPStatus 根据业务错误码获取HTTP状态码
func (receiver BaseResponse) getHTTPStatus(code int) int {
	// 如果是标准HTTP状态码，直接返回
	if code >= 400 && code < 600 {
		return code
	}
	// 其他情况返回200，将业务错误码放在JSON中
	return http.StatusOK
}
