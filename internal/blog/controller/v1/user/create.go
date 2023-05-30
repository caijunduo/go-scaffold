package user

import (
	"github.com/caijunduo/go-scaffold/internal/blog/api/v1"
	"github.com/caijunduo/go-scaffold/internal/blog/biz"
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/caijunduo/go-scaffold/internal/pkg/http"
	"github.com/gin-gonic/gin"
)

func (Controller) Create(c *gin.Context) {
	var req v1.CreateUserRequest
	if err := req.Validate(c); err != nil {
		http.Response(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := biz.User().Create(&req); err != nil {
		http.Response(c, err, nil)
		return
	}

	http.Response(c, nil, nil)
}
