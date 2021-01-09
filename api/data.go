package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAdminTableList(c *gin.Context) {
	fmt.Println("GetAdminTableList")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
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
