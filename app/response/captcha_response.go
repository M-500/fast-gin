package response

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 11:08
//

type CaptchaResponse struct {
	CaptchaID string `json:"captcha_id"`
	PicPath   string `json:"pic_path"`
}
