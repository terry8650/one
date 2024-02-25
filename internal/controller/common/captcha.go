package common

import (
	"context"
	"one2.3/api/v1/common"
	"one2.3/internal/service"
)

var Captcha = cCaptcha{}

type cCaptcha struct{}

func (c *cCaptcha) Get(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().NewAndStore(ctx)
	res = &common.CaptchaRes{
		Key: idKeyC,
		Img: base64stringC,
	}

	return
}
