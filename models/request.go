package models

import (
	"errors"
	"strconv"
)

type PageInfo struct {
	Table       string `json:"table"`
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	PageStr     string `json:"-"`
	PageSizeStr string `json:"-"`
}

func (p *PageInfo) Verify() (err error) {
	if p.Table == "" {
		return errors.New("表名不能为空")
	}
	page, err := strconv.Atoi(p.PageStr)
	if err != nil {
		p.Page = 1
	} else {
		if page < 1 {
			p.Page = 1
		} else {
			p.Page = page
		}
	}
	PageSize, err := strconv.Atoi(p.PageSizeStr)
	if err != nil {
		p.PageSize = 20
	} else {
		if PageSize < 1 {
			p.PageSize = 20
		} else {
			p.PageSize = PageSize
		}
	}
	return nil
}
