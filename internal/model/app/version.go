package app

import (
	"go-admin-x/internal/model/base"
	"time"

	"github.com/jinzhu/gorm"
)

type Version struct {
	base.Base
	Name     string `gorm:"column:name;size:64;" json:"name" form:"name"`
	Platform string `gorm:"column:platform;size:64;not null; "json:"platform" form:"platform"`
	Version  string `gorm:"column:version;size:32;not null;" json:"version" form:"version"`
	Detail   string `gorm:"column:detail;size:4095;" json:"detail" form:"detail"`
	Url      string `gorm:"column:url;size:2047;not null;" json:"url" form:"url"`
	Md5      string `gorm:"column:md5;type:char(32);" json:"md5" form:"md5"`
	Force    uint8  `gorm:"column:force;type:tinyint(1);not null;" json:"force" form:"force"`    // 强制更新(1:强制 0:非强制)
	Status   uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"` // 状态(1:启用 0:不启用)
}

// 表名
func (Version) TableName() string {
	return TableName("version")
}

// 添加前
func (m *Version) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *Version) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}
