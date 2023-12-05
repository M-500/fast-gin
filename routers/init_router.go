package routers

import (
	"fast-gin/app/api"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:20
//

func SetUpRouters() *gin.Engine {
	r := gin.Default()
	naApis := api.APIGroupsAPP.NaApiGroups
	naApis.RegisterRouter(r)
	r.Use()
	adminApis := api.APIGroupsAPP.AdminApiGroups
	adminApis.RegisterRouter(r)
	return r
}
