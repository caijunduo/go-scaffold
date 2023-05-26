package user

import (
	"context"
	"github.com/caijunduo/go-scaffold/internal/blog/model"
	"gorm.io/gorm"
)

type Store interface {
	DB(db *gorm.DB) *store
	Context(ctx context.Context) *store
	Create(model *model.User) error
}

func New() Store {
	return &store{}
}

var _ Store = (*store)(nil)

type store struct {
	db  *gorm.DB
	ctx context.Context
}

func (s *store) DB(db *gorm.DB) *store {
	s.db = db
	return s
}

func (s *store) Context(ctx context.Context) *store {
	s.ctx = ctx
	return s
}

func (s *store) Create(model *model.User) error {
	return s.db.Create(&model).Error
}
