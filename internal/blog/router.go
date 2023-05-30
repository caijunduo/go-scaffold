package blog

import (
	userV1 "github.com/caijunduo/go-scaffold/internal/blog/controller/v1/user"
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/caijunduo/go-scaffold/internal/pkg/http"
	"github.com/caijunduo/go-scaffold/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// installRouters 安装接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		http.Response(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /health handler.
	g.GET("/health", func(c *gin.Context) {
		log.C(c).Info("Health function called")
		http.Response(c, nil, map[string]string{"status": "ok"})
	})

	v1 := g.Group("/v1")
	{
		user := userV1.Controller{}
		// 创建 users 路由分组
		users := v1.Group("/users")
		{
			users.POST("", user.Create)       // 创建用户
			users.GET(":username", user.Show) // 获取用户详情
			users.GET("", user.Index)
		}
	}

	return nil
}
