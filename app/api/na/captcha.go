package na

import (
	"fast-gin/app/service"
	"fast-gin/pkg/comm/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 11:06
var store = base64Captcha.DefaultMemStore

func (*LoginApi) CaptchaImgAPI(ctx *gin.Context) {
	temp, err := service.CaptchaSer.MakeCaptcha()
	if err != nil {
		response.JsonFailMsg(ctx, "生成验证码错误")
		return
	}
	// 将验证码扔进redis中保存
	response.JsonSuccessData(ctx, "验证码生成成功", temp)
	return
}
