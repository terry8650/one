package sys

import (
	"context"
	"one2.3/api/v1/sys"
	"one2.3/internal/consts"
	"one2.3/internal/service"
	"one2.3/utility/response"
)

type cMenu struct{}

var Menu = cMenu{}

func (c *cMenu) Add(ctx context.Context, req *sys.MenuAddReq) (res *sys.MenuAddRes, err error) {

	err = service.Menu().Add(ctx, req)
	if err == nil {
		// 删除相关缓存
		service.Cache().Remove(ctx, consts.SysNode)
	}

	return
}
func (c *cMenu) Edit(ctx context.Context, req *sys.MenuEditReq) (res *sys.MenuEditRes, err error) {
	err = service.Menu().Edit(ctx, req)
	if err == nil {
		// 删除相关缓存
		service.Cache().Remove(ctx, consts.SysNode)
	}

	return
}
func (c *cMenu) Delete(ctx context.Context, req *sys.MenuDeleteReq) (res *sys.MenuDeleteRes, err error) {
	err = service.Menu().Delete(ctx, req.Id)
	if err == nil {
		// 删除相关缓存
		service.Cache().Remove(ctx, consts.SysNode)
	}

	return
}
func (c *cMenu) List(ctx context.Context, req *sys.MenuSearchReq) (res *sys.MenuSearchRes, err error) {

	res, err = service.Menu().GetMenuList(ctx, req)
	if err == nil {
		response.JsonOri(ctx, res)
	}
	return

}
func (c *cMenu) GetMyMenu(ctx context.Context, req *sys.MenuAuthListReq) (res *sys.MenuAuthListRes, err error) {
	res = &sys.MenuAuthListRes{}
	res.AuthList, err = service.Menu().GetAuthList(ctx)

	return
}
