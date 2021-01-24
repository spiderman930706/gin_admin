package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
	"github.com/spiderman930706/gin_admin/service"
	"github.com/spiderman930706/gin_admin/util"
)

func GetAdminTableList(c *gin.Context) {
	var result []string
	for k := range global.Tables {
		result = append(result, k)
	}
	OkWithDetailed(result, "获取成功", c)
}

func GetAdminDataList(c *gin.Context) {
	pageInfo := &models.PageInfo{
		Table:       c.Param("table"),
		PageStr:     c.Query("page"),
		PageSizeStr: c.Query("size"),
	}
	if err := pageInfo.Verify(); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if err, pageResult := service.GetTableDataList(pageInfo); err != nil {
		FailWithMessage("获取失败", c)
	} else {
		OkWithDetailed(pageResult, "获取成功", c)
	}
}

func GetAdminDataDetail(c *gin.Context) {
	dataInfo := &models.DataInfo{
		Table:     c.Param("table"),
		DataIdStr: c.Param("data_id"),
	}
	if err := dataInfo.Verify(true); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if err, dataResult := service.GetDataDetail(dataInfo); err != nil {
		FailWithMessage("获取失败", c)
	} else {
		OkWithDetailed(dataResult, "获取成功", c)
	}
}

func NewAdminData(c *gin.Context) {
	dataInfo := &models.DataInfo{
		Table: c.Param("table"),
	}
	if err := dataInfo.Verify(false); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if !global.Tables[dataInfo.Table].CanAdd {
		FailWithMessage("禁止新增数据", c)
		return
	}
	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		FailWithMessage("请求参数有误", c)
		return
	}
	dataInfo.Data = data
	if err := util.CheckAndChangeData(dataInfo, false); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateData(dataInfo); err != nil {
		FailWithMessage("新增失败，请检查数据", c)
		return
	} else {
		OkWithMessage("新增成功", c)
		return
	}
}

func ChangeAdminData(c *gin.Context) {
	dataInfo := &models.DataInfo{
		Table:     c.Param("table"),
		DataIdStr: c.Param("data_id"),
	}
	if err := dataInfo.Verify(true); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if !global.Tables[dataInfo.Table].CanModify {
		FailWithMessage("禁止修改数据", c)
		return
	}
	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		FailWithMessage("请求参数有误", c)
		return
	}
	dataInfo.Data = data
	if err := util.CheckAndChangeData(dataInfo, true); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if err := service.ChangeData(dataInfo); err != nil {
		FailWithMessage("修改失败", c)
	} else {
		OkWithMessage("修改成功", c)
	}
}

func DeleteAdminData(c *gin.Context) {
	dataInfo := &models.DataInfo{
		Table:     c.Param("table"),
		DataIdStr: c.Param("data_id"),
	}
	if err := dataInfo.Verify(true); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if !global.Tables[dataInfo.Table].CanDelete {
		FailWithMessage("禁止删除数据", c)
		return
	}
	if err := service.DeleteData(dataInfo); err != nil {
		FailWithMessage("删除失败", c)
	} else {
		OkWithMessage("删除成功", c)
	}
}

func BatchDeleteAdminData(c *gin.Context) {
	dataInfo := &models.DataInfo{
		Table: c.Param("table"),
	}
	var idList models.BatchID
	if err := c.ShouldBindJSON(idList); err != nil {
		FailWithMessage("请求参数有误", c)
		return
	}
	if len(idList.IDList) == 0 {
		FailWithMessage("需要携带id_list", c)
		return
	}
	if err := dataInfo.Verify(false); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if !global.Tables[dataInfo.Table].CanDelete {
		FailWithMessage("禁止删除数据", c)
		return
	}
	if err := service.BatchDeleteData(dataInfo, &idList); err != nil {
		FailWithMessage("删除失败", c)
	} else {
		OkWithMessage("删除成功", c)
	}
}
