package models

import (
	"errors"
	"github.com/spiderman930706/gin_admin/global"
	"strconv"
)

type PageInfo struct {
	Table       string `json:"table"`
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	PageStr     string `json:"-"`
	PageSizeStr string `json:"-"`
}

type DataInfo struct {
	Table     string `json:"table"`
	DataId    int    `json:"data_id"`
	DataIdStr string `json:"-"`
}

func (p *PageInfo) Verify() (err error) {
	if err := tableVerify(p.Table); err != nil {
		return err
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

func (d *DataInfo) Verify() (err error) {
	if err := tableVerify(d.Table); err != nil {
		return err
	}
	id, err := strconv.Atoi(d.DataIdStr)
	if err != nil {
		return errors.New("数据id应为数字")
	}
	if id <= 0 {
		return errors.New("数据id过于奇怪")
	}
	d.DataId = id
	return nil
}

func tableVerify(table string) error {
	if table == "" {
		return errors.New("表名不能为空")
	}
	if _, ok := global.Tables[table]; !ok {
		return errors.New("表不存在")
	}
	return nil
}
