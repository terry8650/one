package sys

import (
	"context"
	"one2.3/api/v1/sys"
	"one2.3/internal/service"
	"one2.3/utility/response"
)

type cUser struct{}

var User = cUser{}

func (c *cUser) List(ctx context.Context, req *sys.UserSearchReq) (res *sys.UserSearchRes, err error) {
	res, err = service.User().List(ctx, req)
	if err == nil {
		response.JsonOri(ctx, res)
	}
	return
}
func (c *cUser) GetDeptPostRole(ctx context.Context, req *sys.UserGetDeptPostReq) (res *sys.UserGetDeptPostRes, err error) {
	res = &sys.UserGetDeptPostRes{}
	var dept *sys.DeptListRes
	dept, err = service.Dept().List(ctx)
	if err != nil {
		return nil, err
	}
	res.DeptList = dept.DeptList
	res.PostList, err = service.Post().GetPost(ctx)
	if err != nil {
		return nil, err
	}
	res.RoleList, err = service.Role().GetRole(ctx)
	return
}
func (c *cUser) Add(ctx context.Context, req *sys.UserAddReq) (res *sys.UserAddRes, err error) {
	err = service.User().Add(ctx, req)
	return
}
func (c *cUser) Update(ctx context.Context, req *sys.UserUpdateReq) (res *sys.UserUpdateRes, err error) {
	err = service.User().Update(ctx, req)
	return
}
func (c *cUser) Delete(ctx context.Context, req *sys.UserDeleteReq) (res *sys.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req)
	return
}
func (c *cUser) GetRoleIds(ctx context.Context, req *sys.UserGetRoleReq) (res *sys.UserGetRoleRes, err error) {
	res = &sys.UserGetRoleRes{}
	res.RoleIds, err = service.Cas().GetUserRoles(req.Id)
	return
}
func (c *cUser) ChangePwd(ctx context.Context, req *sys.UserChangePwdReq) (res *sys.UserChangePwdRes, err error) {
	err = service.User().ChangePwd(ctx, req)
	return nil, err
}
