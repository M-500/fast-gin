package na

import "github.com/gin-gonic/gin"

//
// @Description 这里是不需要登录的API集合
// @Author 代码小学生王木木
// @Date 2023/12/5 10:16
//

type NaApis struct {
	LoginApi
}

func (*NaApis) RegisterRouter(r *gin.Engine) {
	group := r.Group("/na")
	na := new(NaApis)
	group.POST("/login", na.LoginHandler)
	group.GET("/captcha", na.CaptchaImgAPI)
}
