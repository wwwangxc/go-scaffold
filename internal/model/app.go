package model

import (
	"go-scaffold/pkg/xgorm"
	"time"

	"github.com/jinzhu/gorm"
)

// App ..
type App struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	AppID     string    `json:"app_id" gorm:"column:app_id" description:"租户id	"`
	Name      string    `json:"name" gorm:"column:name" description:"租户名称	"`
	Secret    string    `json:"secret" gorm:"column:secret" description:"密钥"`
	WhiteIPS  string    `json:"white_ips" gorm:"column:white_ips" description:"ip白名单，支持前缀匹配"`
	QPD       int64     `json:"qpd" gorm:"column:qpd" description:"日请求量限制"`
	QPS       int64     `json:"qps" gorm:"column:qps" description:"每秒请求量限制"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间	"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	IsDelete  bool      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

// TableName ..
func (t *App) TableName() string {
	return "app"
}

// GetByID 通过id获取app信息
func (t *App) GetByID(id int64) error {
	if err := xgorm.Cli.Where("id = ?", id).Find(t).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
