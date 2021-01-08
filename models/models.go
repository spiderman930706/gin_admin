package models

import "time"

type AdminOperation interface {
	CanDelete()
	CanModify()
}

type Model struct {
	ID         int `gorm:"primary_key" json:"id" list:"id" type:"int"`
	CreatedOn  time.Time `json:"create_time" list:"create_time" type:"time"`
	ModifiedOn time.Time `json:"modified_time" list:"modified_time" type:"time"`
}

func (m Model) CanDelete() bool {
	return true
}

func (m Model) CanModify() bool {
	return true
}

