package model

import (
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: 1999single
 * @Description:
 * @File: base
 * @Version: 1.0.0
 * @Date: 2022/5/10 23:29
 */
type Base struct {
	ID        int64 `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
