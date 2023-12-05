package admin

import "github.com/gin-gonic/gin"

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:16

type AdminApis struct {
	UserApi
}

func (*AdminApis) RegisterRouter(r *gin.Engine) {
	group := r.Group("/admin")
	na := new(AdminApis)
	group.GET("/user-info", na.UserInfoHandler)
}
