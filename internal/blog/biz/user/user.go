package user

import (
	"context"
	"errors"
	"github.com/caijunduo/go-scaffold/internal/blog/api/v1"
	"github.com/caijunduo/go-scaffold/internal/blog/model"
	"github.com/caijunduo/go-scaffold/internal/blog/store"
	"github.com/caijunduo/go-scaffold/internal/pkg/errno"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"regexp"
)

type Biz interface {
	Context(ctx context.Context) *biz
	Create(req *v1.CreateUserRequest) error
	Show(username string) (*v1.ShowUserResponse, error)
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

func (b *biz) Create(req *v1.CreateUserRequest) error {
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

func (b *biz) Show(username string) (*v1.ShowUserResponse, error) {
	user, err := store.User().GetUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}

	var resp v1.ShowUserResponse
	_ = copier.Copy(&resp, user)

	return &resp, nil
}
