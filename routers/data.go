package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/spiderman930706/gin_admin/api"
)

func InitTableRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("tables")
	{
		apiRouter.GET("/list", api.GetAdminTableList)
		apiRouter.GET("/dataList/:table", api.GetAdminDataList)
		apiRouter.GET("/dataDetail/:table/:data_id", api.GetAdminDataDetail)
		apiRouter.POST("/newData/:table", api.NewAdminData)
		apiRouter.PUT("/changeData/:table/:data_id", api.ChangeAdminData)
		apiRouter.DELETE("/deleteData/:table/:data_id", api.DeleteAdminData)
	}
}
