package service

import (
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
)

func GetTableDataList(info models.PageInfo) (err error, list interface{}, dict map[string]map[string]interface{}, total int64) {
	limit := info.PageSize
	table := info.Table
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Table(table)
	var result []map[string]interface{}
	err = db.Count(&total).Error
	selectName, dict := listSelectName(table)
	err = db.Select(selectName).Limit(limit).Offset(offset).Find(&result).Error
	list = filterListData(table, result)
	return err, list, dict, total
}

//获取数据列表页字段字典和查询字段名
func listSelectName(table string) (list []string, dict map[string]map[string]interface{}) {
	tableInfo := global.Tables[table]
	dict = make(map[string]map[string]interface{})
	fields := tableInfo.Field
	for k, v := range fields {
		if v.List != "" && v.Type != "password" {
			var typeName interface{}
			if v.Type == "" {
				typeName = v.Schema.GORMDataType
			} else {
				typeName = v.Type
			}
			dataDict := make(map[string]interface{})
			dataDict["name"] = v.List
			dataDict["type"] = typeName
			dict[k] = dataDict
			list = append(list, k)
		}
	}
	return
}

//对列表展示页结果数据进行处理
func filterListData(table string, result []map[string]interface{}) interface{} {
	tableInfo := global.Tables[table]
	fields := tableInfo.Field
	for _, data := range result {
		for k := range data {
			info := fields[k]
			switch info.Type {
			case "password":
				data[k] = ""
			case "time":
				//todo 处理时间转换
			}
		}
	}
	return result
}
