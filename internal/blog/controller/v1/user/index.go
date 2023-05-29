package user

import (
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/caijunduo/go-scaffold/internal/pkg/http"
	"github.com/caijunduo/go-scaffold/pkg/api/blog_api_v1"
	"github.com/gin-gonic/gin"
)

func (Controller) Index(c *gin.Context) {
	var req blog_api_v1.GetUserRequest
	if err := req.Validate(c); err != nil {
		http.Response(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}
	http.Response(c, nil, req)
}
