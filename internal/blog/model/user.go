package model

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	Username  string    `gorm:"column:username"`       //
	Password  string    `gorm:"column:password"`       //
	Email     string    `gorm:"column:email"`          //
	Phone     string    `gorm:"column:phone"`          //
	CreatedAt time.Time `gorm:"column:created_at"`     //
	UpdatedAt time.Time `gorm:"column:updated_at"`     //
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "`blog`.`users`"
}
