package na

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:17
//

type LoginApi struct {
}

func (*LoginApi) LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
