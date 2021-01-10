package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/spiderman930706/gin_admin/global"
)

func GetAdminTableList(c *gin.Context) {
	var result []string
	for k := range global.Tables {
		result = append(result, k)
	}
	OkWithData(result, c)
}

func GetAdminTableData(c *gin.Context) {
	//result := map[string]interface{}{}
	//global.DB.Table("users").Take(&result)

	table := c.Param("table")
	fmt.Println(table)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func GetAdminDataDetail(c *gin.Context) {
	table := c.Param("table")
	fmt.Println(table)
	dataId := c.Param("data_id")
	fmt.Println(dataId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func NewAdminData(c *gin.Context) {
	table := c.Param("table")
	fmt.Println(table)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func ChangeAdminData(c *gin.Context) {
	table := c.Param("table")
	fmt.Println(table)
	dataId := c.Param("data_id")
	fmt.Println(dataId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func DeleteAdminData(c *gin.Context) {
	table := c.Param("table")
	fmt.Println(table)
	dataId := c.Param("data_id")
	fmt.Println(dataId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
