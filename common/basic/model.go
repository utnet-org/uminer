package dao

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// 数据表的struct需要包含Model
type Model struct {
	CreatedAt time.Time             `gorm:"type:datetime(3);not null"`
	UpdatedAt time.Time             `gorm:"type:datetime(3);not null"`
	DeletedAt soft_delete.DeletedAt `gorm:"index"`
}
