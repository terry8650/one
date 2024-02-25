package captcha

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
	"one2.3/internal/service"
)

type sCaptcha struct{}

var (
	captchaStore  = base64Captcha.DefaultMemStore
	captchaDriver = newDriver()
)

func init() {
	service.RegisterCaptcha(New())
}

// Captcha 验证码管理服务
func New() *sCaptcha {
	return &sCaptcha{}
}

func newDriver() *base64Captcha.DriverString {
	driver := &base64Captcha.DriverString{
		Height:          44,
		Width:           126,
		NoiseCount:      5,
		ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "1234567890",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

// NewAndStore 创建验证码，直接输出验证码图片内容到HTTP Response.
func (s *sCaptcha) NewAndStore(ctx context.Context) (idKeyC string, base64stringC string, err error) {
	var (
		c = base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	)
	idKeyC, base64stringC, err = c.Generate()
	return
}

// VerifyAndClear 校验验证码
func (s *sCaptcha) Verify(id, answer string) bool {
	answer = gstr.ToLower(answer)
	return captchaStore.Verify(id, answer, true)
}
