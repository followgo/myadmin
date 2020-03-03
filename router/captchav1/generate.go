package captchav1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mojocn/base64Captcha"
)

// GenerateHandler 生产验证码
func GenerateHandler(c echo.Context) error {
	config := new(configJsonBody)
	if err := c.Bind(config); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}
	if config.ImgWidth <= 0 {
		config.ImgWidth = 240
	}
	if config.ImgHeight <= 0 {
		config.ImgHeight = 60
	}

	var driver base64Captcha.Driver

	// create base64 encoding captcha
	switch config.CaptchaType {
	case "audio":
		driver = &base64Captcha.DriverAudio{
			Length:   4,
			Language: "en",
		}

	case "string":
		driverStr := &base64Captcha.DriverString{
			Height: config.ImgHeight, Width: config.ImgWidth,
			ShowLineOptions: base64Captcha.OptionShowHollowLine,
			Source:          "1234567890QWERTYUIOPLKJHGFDSAZXCVBNM",
			Length:          4,
			Fonts:           fonts,
		}
		driverStr.ConvertFonts()
		driver = driverStr

	case "math":
		driverMath := &base64Captcha.DriverMath{
			Height: config.ImgHeight, Width: config.ImgWidth,
			ShowLineOptions: base64Captcha.OptionShowHollowLine,
			Fonts:           fonts,
		}
		driverMath.ConvertFonts()
		driver = driverMath

	// case "chinese":
	// 	driverChinese = &base64Captcha.DriverChinese{
	// 		Height: config.ImgHeight, Width: config.ImgWidth,
	// 		ShowLineOptions: base64Captcha.OptionShowHollowLine,
	// 		Length:          2,
	// 		Source:          base64Captcha.TxtChineseCharaters,
	// 		Fonts:           []string{"wqy-microhei.ttc"}, // 中文字体
	// 	}
	// 	driverChinese.ConvertFonts()
	// 	driver = driverChinese

	default:
		driverDigit := &base64Captcha.DriverDigit{
			Height: config.ImgHeight, Width: config.ImgWidth,
			Length:  4,
			MaxSkew: 0.7, DotCount: 80,
		}
		driver = driverDigit
	}

	ca := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := ca.Generate()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "生产验证码", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"data": b64s, "captcha_id": id})
}
