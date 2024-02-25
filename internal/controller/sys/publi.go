package sys

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/api/v1/sys"
	"one2.3/internal/model"
	"one2.3/internal/service"
)

type cPubli struct{}

var Publi = cPubli{}

func (c *cPubli) Login(ctx context.Context, req *sys.UserLoginReq) (res *sys.UserLoginRes, err error) {

	res, err = service.Gftoken().Login(ctx, &model.UserLoginInput{
		Username: req.Username,
		Password: req.Password,
		Captcha:  req.Captcha,
		IdKey:    req.IdKey,
	})

	return
}
func (c *cPubli) Tes(ctx context.Context, req *sys.TestReq) (res *sys.TestRes, err error) {
	r := g.RequestFromCtx(ctx)
	g.Dump(r.GetBody())

	return
}
