package service

import (
	"fmt"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
)

func GetTableDataList(info models.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	table := info.Table
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Table(table)
	tableInfo := global.Tables[table]
	fmt.Println(tableInfo)
	result := map[string]interface{}{}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Take(&result).Error
	return err, result, total
}
