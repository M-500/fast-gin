package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:18
//

type UserApi struct {
}

func (*UserApi) UserInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "用户列表",
	})
}
