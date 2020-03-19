package sys

import (
	"go-admin-x/internal/model/base"
	"go-admin-x/internal/util/orm"
	"time"

	"github.com/jinzhu/gorm"
)

// 用户-角色
type AdminRole struct {
	base.Base
	AdminID uint64 `gorm:"column:admin_id;unique_index:uk_admin_role_admin_id;not null;"` // 管理员ID
	RoleID  uint64 `gorm:"column:role_id;unique_index:uk_admin_role_admin_id;not null;"`  // 角色ID
}

// 表名
func (AdminRole) TableName() string {
	return TableName("admin_role")
}

// 添加前
func (m *AdminRole) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// 更新前
func (m *AdminRole) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now()
	return nil
}

// 分配用户角色
func (AdminRole) SetRole(adminsid uint64, roleids []uint64) error {
	tx := orm.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(&AdminRole{AdminID: adminsid}).Delete(&AdminRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(roleids) > 0 {
		for _, rid := range roleids {
			rm := new(AdminRole)
			rm.RoleID = rid
			rm.AdminID = adminsid
			if err := tx.Create(rm).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
