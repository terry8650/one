package home

import (
	"context"
	"one2.3/api/v1/home"
)

type cWechat struct{}

var Wechat = cWechat{}

func (c *cWechat) Login(ctx context.Context, req *home.WechatLoginReq) (res *home.WechatLoginRes, err error) {
	return
}
func (c *cWechat) CallBack(ctx context.Context, req *home.CallBackReq) (res *home.CallBackRes, err error) {
	return
}
