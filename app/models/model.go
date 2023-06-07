package models

import (
	"github.com/spf13/cast"
	"time"
)

// 模型通用属性和方法
type BaseModel struct {
	ID uint64 `colunm:id;primaryKey;autoIncrement;
json:"id,omitempty"`
}

// 时间戳
type CommonTimestampsField struct {
	CreateAt time.Time `gorm:"column:created_at;index;"
json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;"
json:"updated_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
