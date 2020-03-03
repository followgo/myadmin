package captchav1

import (
	"github.com/mojocn/base64Captcha"
)

// configJsonBody json request body.
type configJsonBody struct {
	Id            string `json:"id"`
	CaptchaType   string `json:"captcha_type"`
	VerifyValue   string `json:"verify_value"`
	ImgWidth      int    `json:"img_width"`
	ImgHeight     int    `json:"img_height"`
}

// store 使用自带memory store
var store = base64Captcha.DefaultMemStore


// fonts 验证码字符的字体
var fonts = []string{
	"./router/captchav1/fonts/3Dumb.ttf",
	"./router/captchav1/fonts/actionj.ttf",
	"./router/captchav1/fonts/Flim-Flam.ttf",
	"./router/captchav1/fonts/ApothecaryFont.ttf",
	"./router/captchav1/fonts/DENNEthree-dee.ttf",
	"./router/captchav1/fonts/DeborahFancyDress.ttf",
	"./router/captchav1/fonts/chromohv.ttf",
}
