package admin

import (
	"fast-gin/pkg/comm/response"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:18
//

type UserApi struct {
}

func (*UserApi) UserInfoHandler(c *gin.Context) {
	response.JsonFailData(c, "成功获取用户", "傻逼")
}
