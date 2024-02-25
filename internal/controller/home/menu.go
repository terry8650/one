package home

import (
	"context"
	"one2.3/api/v1/home"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

func (c *cMenu) GetMenu(ctx context.Context, req *home.HMenuReq) (res *home.HMenuRes, err error) {
	res = &home.HMenuRes{}
	res.AuthList = []string{}
	return
}
