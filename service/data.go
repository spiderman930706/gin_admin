package service

import (
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
)

//列表展示页数据获取
func GetTableDataList(info models.PageInfo) (err error, pageResult models.PageResult) {
	limit := info.PageSize
	table := info.Table
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Table(table)
	var result []map[string]interface{}
	var list interface{}
	var total int64
	err = db.Count(&total).Error
	selectName, dict := listSelectName(table)
	err = db.Select(selectName).Limit(limit).Offset(offset).Find(&result).Error
	list = filterListData(table, result)
	pageResult = models.PageResult{
		Items:     list,
		Dict:      dict,
		Total:     total,
		Page:      info.Page,
		PageSize:  info.PageSize,
		CanModify: global.Tables[table].CanModify,
		CanDelete: global.Tables[table].CanDelete,
	}
	return err, pageResult
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

//根据id获取数据
func GetDataDetail(info models.DataInfo) (err error, pageResult models.DataResult) {
	table := info.Table
	var result = make(map[string]interface{})
	dict := DataMap(table)

	err = global.DB.Table(table).Where("id = ?", info.DataId).Take(&result).Error
	pageResult = models.DataResult{
		Item:      result,
		Dict:      dict,
		CanModify: global.Tables[table].CanModify,
		CanDelete: global.Tables[table].CanDelete,
	}
	return err, pageResult
}

func DataMap(table string) (dict map[string]map[string]interface{}) {
	tableInfo := global.Tables[table]
	dict = make(map[string]map[string]interface{})
	fields := tableInfo.Field
	for k, v := range fields {
		var typeName interface{}
		if v.Type == "" {
			typeName = v.Schema.GORMDataType
		} else {
			typeName = v.Type
		}
		dataDict := make(map[string]interface{})
		if v.List != "" {
			dataDict["name"] = v.List
		} else {
			dataDict["name"] = k
		}
		dataDict["type"] = typeName
		dict[k] = dataDict
	}
	return
}
