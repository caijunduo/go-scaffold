package user

import (
	"github.com/caijunduo/go-scaffold/internal/blog/biz"
	"github.com/caijunduo/go-scaffold/internal/pkg/http"
	"github.com/caijunduo/go-scaffold/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (Controller) Show(c *gin.Context) {
	log.C(c).Info("Create user function called")

	user, err := biz.User().Show(c.Param("username"))
	if err != nil {
		http.Response(c, err, nil)
		return
	}

	http.Response(c, nil, user)
}
