package user

import (
	"context"
	"github.com/caijunduo/go-scaffold/internal/blog/model"
	"github.com/caijunduo/go-scaffold/internal/blog/store"
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/caijunduo/go-scaffold/pkg/api/blog_api_v1"
	"github.com/jinzhu/copier"
	"regexp"
)

type Biz interface {
	Context(ctx context.Context) *biz
	Create(req *blog_api_v1.CreateUserRequest) error
}

func New() Biz {
	return &biz{}
}

var _ Biz = (*biz)(nil)

type biz struct {
	ctx context.Context
}

func (b *biz) Context(ctx context.Context) *biz {
	b.ctx = ctx
	return b
}

func (b *biz) Create(req *blog_api_v1.CreateUserRequest) error {
	var m model.User
	_ = copier.Copy(&m, req)

	if err := store.User().Create(&m); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}
