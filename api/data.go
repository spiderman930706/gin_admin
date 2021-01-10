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
	if err, list, dict, total := service.GetTableDataList(pageInfo); err != nil {
		FailWithMessage("获取失败", c)
	} else {
		OkWithDetailed(models.PageResult{
			Items:    list,
			Dict:     dict,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func GetAdminDataDetail(c *gin.Context) { // todo
	table := c.Param("table")
	fmt.Println(table)
	dataId := c.Param("data_id")
	fmt.Println(dataId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
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
