package model

import (
	"go-scaffold/internal/constant"
	xgorm "go-scaffold/pkg/database/xgorm"
	"time"

	"github.com/jinzhu/gorm"
)

// Admin table user
type Admin struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"user_name" gorm:"column:user_name"`
	Salt      string    `json:"salt" gorm:"column:salt"`
	Password  string    `json:"password" gorm:"column:password"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete"`
}

// TableName ..
func (t *Admin) TableName() string {
	return "admin"
}

// GetByUsername ..
func (t *Admin) GetByUsername(username string) error {
	if err := xgorm.Store(constant.MySQLStoreNameDB1).
		Where("user_name = ?", username).Find(t).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	return nil
}
