package model

import (
	"context"
	"fmt"
	"go-scaffold/internal/constant"

	"github.com/jinzhu/gorm"
)

// User table user
type User struct {
	gorm.Model

	UserName string `json:"user_name" gorm:"column:user_name"`
	Salt     string `json:"salt" gorm:"column:salt"`
	Password string `json:"password" gorm:"column:password"`
}

// TableName ..
func (u *User) TableName() string {
	return "admin"
}

// FindUserByID get user by id
func FindUserByID(ctx context.Context, id int) (*User, error) {
	return FindUserWithOption(ctx, WithIDEqual(id))
}

// FindUserByUserName find user by user name
func FindUserByUserName(ctx context.Context, userName string) (*User, error) {
	return FindUserWithOption(ctx, WithUserNameEuqa(userName))
}

// FindUserWithOption find user with options
func FindUserWithOption(ctx context.Context, opts ...Option) (*User, error) {
	builder, err := newSQLBuilder(constant.MySQLStoreNameMaster, opts...)
	if err != nil {
		return nil, fmt.Errorf("new sql builder fail. err:%w", err)
	}

	db, err := builder.build(ctx)
	if err != nil {
		return nil, fmt.Errorf("sql builder build fail. err:%w", err)
	}

	user := &User{}
	ret := db.First(user)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return user, nil
}
