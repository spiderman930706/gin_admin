package util

import (
	"errors"
	"fmt"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
	"time"
)

//获取所有字段字典
func DataMap(table string) (dict map[string]*models.Dict) {
	tableInfo := global.Tables[table]
	dict = make(map[string]*models.Dict)
	fields := tableInfo.Field
	for k, v := range fields {
		var typeName string
		if v.Type == "" {
			typeName = string(v.Schema.GORMDataType)
		} else {
			typeName = v.Type
		}
		dataDict := &models.Dict{}
		if v.Name != "" {
			dataDict.Name = v.Name
		} else {
			dataDict.Name = k
		}
		dataDict.List = v.List
		dataDict.Type = typeName
		dict[k] = dataDict
	}
	return
}

func CheckAndChangeData(info *models.DataInfo, modify bool) (err error) {
	UpdateTime(modify, info.Data)
	if err := CheckFields(info); err != nil {
		return err
	}
	return
}

func UpdateTime(modify bool, data map[string]interface{}) {
	if !modify {
		data["created_on"] = time.Now()
	}
	data["modified_on"] = time.Now()
}

func CheckFields(info *models.DataInfo) error {
	fields := global.Tables[info.Table].Field
	for k, v := range info.Data {
		if k == "id" {
			delete(info.Data, k)
		}
		if table, ok := fields[k]; ok {
			dataType := table.Type
			switch dataType {
			case "password":
				if value, ok := v.(string); ok {
					info.Data[k] = MD5V([]byte(value))
				} else {
					return errors.New("密码居然不是字符串？")
				}
			case "time":
				//todo 时间类型转换
				fmt.Println("时间")
			}
		} else {
			return errors.New(fmt.Sprintf("%s 字段不存在", k))
		}
	}
	return nil
}
