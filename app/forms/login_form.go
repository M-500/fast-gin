package forms

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 11:05
//

// PwdLoginForm
// @Description: 用户账号密码登录
type PwdLoginForm struct {
	UserName    string `form:"username" json:"username" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required,min=6,max=30"`
	CaptchaCode string `form:"captcha_code" json:"captcha_code" binding:"required,min=3,max=6"`
	CaptchaId   string `form:"captcha_id" json:"captcha_id" binding:"required,min=3,max=200"`
	LoginType   string `form:"login_type" json:"login_type" binding:"required,oneof=cms wechat"`
}
