package models

import (
	"errors"
	"strconv"

	"github.com/spiderman930706/gin_admin/global"
)

type TableInfo struct {
	Table string
}

type PageInfo struct {
	Table       string
	Page        int
	PageSize    int
	PageStr     string
	PageSizeStr string
}

type DataInfo struct {
	Table     string
	DataId    int
	DataIdStr string
	Data      map[string]interface{}
}

type BatchID struct {
	IDList []int `json:"id_list"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (t *TableInfo) Verify() (err error) {
	if err := tableVerify(t.Table); err != nil {
		return err
	}
	return
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

func (d *DataInfo) Verify(checkId bool) (err error) {
	if err := tableVerify(d.Table); err != nil {
		return err
	}
	if checkId {
		id, err := strconv.Atoi(d.DataIdStr)
		if err != nil {
			return errors.New("数据id应为数字")
		}
		if id <= 0 {
			return errors.New("数据id过于奇怪")
		}
		d.DataId = id
	}
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

func (l *Login) Verify() (err error) {
	if l.Password == "" || l.Username == "" {
		return errors.New("缺少请求参数")
	}
	return
}
