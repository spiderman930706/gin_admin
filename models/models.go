package models

import "time"

type AdminOperation interface {
	CanDelete() bool
	CanModify() bool
}

type Model struct {
	ID         uint      `gorm:"primary_key" json:"id" admin:"list:id;type:int"`
	CreatedOn  time.Time `json:"create_time" admin:"list:create_time;type:time"`
	ModifiedOn time.Time `json:"modified_time" admin:"list:modified_time;type:time"`
}

//能否删除字段
func (m Model) CanDelete() bool {
	return true
}

//能否修改字段
func (m Model) CanModify() bool {
	return true
}
