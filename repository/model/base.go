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
	ID        int64          `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
