package routers

import (
	"fast-gin/app/api"
	"fast-gin/middlewares"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:20
//

func SetUpRouters() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.PaginationMiddleware()) // 注册分页逻辑
	naApis := api.APIGroupsAPP.NaApiGroups
	naApis.RegisterRouter(r)
	r.Use()
	adminApis := api.APIGroupsAPP.AdminApiGroups
	adminApis.RegisterRouter(r)
	return r
}
