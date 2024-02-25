package sys

import (
	"context"
	"one2.3/api/v1/sys"
	"one2.3/internal/service"
)

type cConf struct{}

var Conf = cConf{}

func (c *cConf) List(ctx context.Context, req *sys.ConfListReq) (res *sys.ConfListRes, err error) {
	res, err = service.Conf().List(ctx)
	return
}
func (c *cConf) GetVal(ctx context.Context, req *sys.ConfGetReq) (res *sys.ConfGetRes, err error) {
	res, err = service.Conf().GetVal(ctx, req)
	return
}
func (c *cConf) Add(ctx context.Context, req *sys.ConfAddReq) (res *sys.ConfAddRes, err error) {
	err = service.Conf().Add(ctx, req)
	return
}
func (c *cConf) Update(ctx context.Context, req *sys.ConfUpdateReq) (res *sys.ConfUpdateRes, err error) {
	err = service.Conf().Update(ctx, req)
	return
}
func (c *cConf) Delete(ctx context.Context, req *sys.ConfDeleteReq) (res *sys.ConfDeleteRes, err error) {
	err = service.Conf().Delete(ctx, req.Ids)
	return
}
