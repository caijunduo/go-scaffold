package biz

import (
	"github.com/caijunduo/go-scaffold/internal/blog/biz/user"
)

// User 返回一个实现了 Biz 接口的实例.
func User() user.Biz {
	return user.New()
}
