package blog

import (
	"github.com/caijunduo/go-scaffold/internal/blog/controller/v1/user"
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/caijunduo/go-scaffold/internal/pkg/http"
	"github.com/caijunduo/go-scaffold/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		http.Response(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")
		http.Response(c, nil, map[string]string{"status": "ok"})
	})

	//authz, err := auth.NewAuthz(store.S.DB())
	//if err != nil {
	//	return err
	//}

	uc := user.New()
	//
	//g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create) // 创建用户
			//userv1.PUT(":name/change-password", uc.ChangePassword) // 修改用户密码
			//userv1.Use(mw.Authn(), mw.Authz(authz))
			//userv1.GET(":name", uc.Get) // 获取用户详情
		}
	}

	return nil
}
