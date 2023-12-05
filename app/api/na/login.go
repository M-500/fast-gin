package na

import (
	"fast-gin/app/forms"
	"fast-gin/pkg/comm/validators"
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
	loginForm := forms.PwdLoginForm{}
	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		validators.HandlerValidatorError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
