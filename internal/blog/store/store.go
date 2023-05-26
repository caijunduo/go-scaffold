package store

import (
	"github.com/caijunduo/go-scaffold/internal/blog/store/user"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	ds   *datastore
)

// datastore 是 IStore 的一个具体实现.
type datastore struct {
	db *gorm.DB
}

func Init(db *gorm.DB) {
	once.Do(func() {
		ds = &datastore{
			db: db,
		}
	})
}

// DB 返回存储在 datastore 中的 *gorm.DB.
func DB() *gorm.DB {
	return ds.db
}

// User 返回一个实现了 Store 接口的实例.
func User() user.Store {
	return user.New().DB(ds.db)
}
