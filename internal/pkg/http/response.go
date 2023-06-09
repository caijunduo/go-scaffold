package http

import (
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrResponse 定义了发生错误时的返回消息.
type ErrResponse struct {
	// Code 指定了业务错误码.
	Code string `json:"code,omitempty"`

	// Message 包含了可以直接对外展示的错误信息.
	Message string `json:"message,omitempty"`
}

func Response(c *gin.Context, err error, data interface{}) {
	if err != nil {
		status, code, message := errno.Decode(err)
		c.JSON(status, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
