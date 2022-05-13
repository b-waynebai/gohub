// 模型通用属性与方法
package models

import "time"

// BaseModel 模型基类
type BaseModel struct {
	ID int64 `gorm:"column:id;parimaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	Created time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	Updated time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}
