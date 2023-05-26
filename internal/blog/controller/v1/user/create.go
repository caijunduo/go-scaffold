package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/caijunduo/go-scaffold/internal/blog/biz"
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/caijunduo/go-scaffold/internal/pkg/http"
	"github.com/caijunduo/go-scaffold/internal/pkg/log"
	"github.com/caijunduo/go-scaffold/pkg/api/blog_api_v1"
	"github.com/gin-gonic/gin"
)

func (UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var req blog_api_v1.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.Response(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		http.Response(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := biz.User().Create(&req); err != nil {
		http.Response(c, err, nil)
		return
	}

	http.Response(c, nil, nil)
}
