package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
	"github.com/spiderman930706/gin_admin/service"
)

func GetAdminTableList(c *gin.Context) {
	var result []string
	for k := range global.Tables {
		result = append(result, k)
	}
	OkWithDetailed(result, "获取成功", c)
}

func GetAdminDataList(c *gin.Context) {
	pageInfo := models.PageInfo{
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
	dataInfo := models.DataInfo{
		Table:     c.Param("table"),
		DataIdStr: c.Param("data_id"),
	}
	if err := dataInfo.Verify(); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if err, dataResult := service.GetDataDetail(dataInfo); err != nil {
		FailWithMessage("获取失败", c)
	} else {
		OkWithDetailed(dataResult, "获取成功", c)
	}
}

func NewAdminData(c *gin.Context) { // todo
	table := c.Param("table")
	fmt.Println(table)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func ChangeAdminData(c *gin.Context) { // todo
	table := c.Param("table")
	fmt.Println(table)
	dataId := c.Param("data_id")
	fmt.Println(dataId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func DeleteAdminData(c *gin.Context) { // todo
	table := c.Param("table")
	fmt.Println(table)
	dataId := c.Param("data_id")
	fmt.Println(dataId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
