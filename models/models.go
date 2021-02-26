package models

import (
	"time"

	"gorm.io/gorm"
)

type AdminOperation interface {
	CanDelete() bool
	CanModify() bool
	CanAdd() bool
}

type Model struct {
	ID        uint      `gorm:"primary_key" json:"id" admin:"list;type:int;name:序号"`
	CreatedAt time.Time `json:"created_at" admin:"list;type:time;name:添加时间" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" admin:"list;type:time:修改时间" gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
}

//能否删除数据
func (m Model) CanDelete() bool {
	return true
}

//能否修改数据
func (m Model) CanModify() bool {
	return true
}

//能否新增数据
func (m Model) CanAdd() bool {
	return true
}
