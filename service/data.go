package service

import (
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
	"github.com/spiderman930706/gin_admin/util"
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
		CanAdd:    global.Tables[table].CanAdd,
	}
	return err, pageResult
}

//获取数据列表页字段字典和查询字段名
func listSelectName(table string) (list []string, dict map[string]*models.Dict) {
	dict = util.DataMap(table)
	for k, v := range dict {
		if v.List && v.Type != "password" {
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
	dict := util.DataMap(table)

	err = global.DB.Table(table).Where("id = ?", info.DataId).Take(&result).Error
	pageResult = models.DataResult{
		Item:      result,
		Dict:      dict,
		CanModify: global.Tables[table].CanModify,
		CanDelete: global.Tables[table].CanDelete,
	}
	return err, pageResult
}

//新增数据
func CreateData(info models.DataInfo) (err error) {
	data := info.Data
	result := global.DB.Table(info.Table).Create(&data)
	err = result.Error
	return err
}

//修改数据
func ChangeData(info models.DataInfo) (err error) {
	data := info.Data
	result := global.DB.Table(info.Table).Where("id = ?", info.DataId).Updates(data)
	err = result.Error
	return err
}

//删除数据
func DeleteData(info models.DataInfo) (err error) {
	model := global.Tables[info.Table].Source
	result := global.DB.Table(info.Table).Where("id = ?", info.DataId).Delete(&model) //一定要传？其实传什么都可以
	err = result.Error
	return err
}
