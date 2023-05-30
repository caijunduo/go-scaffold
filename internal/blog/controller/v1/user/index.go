package user

import (
	"github.com/caijunduo/go-scaffold/internal/blog/api/v1"
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/caijunduo/go-scaffold/internal/pkg/http"
	"github.com/gin-gonic/gin"
)

func (Controller) Index(c *gin.Context) {
	var req v1.GetUserRequest
	if err := req.Validate(c); err != nil {
		http.Response(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}
	http.Response(c, nil, req)
}
