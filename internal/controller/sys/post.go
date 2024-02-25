package sys

import (
	"context"
	"one2.3/api/v1/sys"
	"one2.3/internal/service"
	"one2.3/utility/response"
)

type cPost struct{}

var Post = cPost{}

func (c *cPost) List(ctx context.Context, req *sys.PostSearchReq) (res *sys.PostSearchRes, err error) {
	res, err = service.Post().List(ctx, req)
	if err == nil {
		response.JsonOri(ctx, res)
	}
	return
}
func (c *cPost) Add(ctx context.Context, req *sys.PostAddReq) (res *sys.PostAddRes, err error) {

	err = service.Post().Add(ctx, req)

	return nil, err
}

// Edit 修改岗位
func (c *cPost) Edit(ctx context.Context, req *sys.PostEditReq) (res *sys.PostEditRes, err error) {
	err = service.Post().Edit(ctx, req)
	return
}

// Delete 删除岗位
func (c *cPost) Delete(ctx context.Context, req *sys.PostDeleteReq) (res *sys.PostDeleteRes, err error) {
	err = service.Post().Delete(ctx, req.Ids)
	return
}
