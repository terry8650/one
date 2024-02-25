package sys

import (
	"context"
	"one2.3/api/v1/sys"
	"one2.3/internal/service"
	"one2.3/utility/response"
)

type cDept struct{}

var Dept = cDept{}

func (c *cDept) List(ctx context.Context, req *sys.DeptListReq) (res *sys.DeptListRes, err error) {

	res, err = service.Dept().List(ctx)
	if err == nil {
		response.JsonOri(ctx, res)
	}

	return
}
func (c *cDept) Add(ctx context.Context, req *sys.DeptAddReq) (res *sys.DeptAddRes, err error) {

	err = service.Dept().Add(ctx, req)

	return nil, err
}

// Edit 修改岗位
func (c *cDept) Edit(ctx context.Context, req *sys.DeptUpdateReq) (res *sys.DeptUpdateRes, err error) {
	err = service.Dept().Update(ctx, req)
	return
}

// Delete 删除岗位
func (c *cDept) Delete(ctx context.Context, req *sys.DeptDeleteReq) (res *sys.DeptDeleteRes, err error) {
	err = service.Dept().Delete(ctx, req.Id)
	return
}
