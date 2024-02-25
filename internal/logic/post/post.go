package post

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"one2.3/api/v1/sys"
	"one2.3/internal/consts"
	"one2.3/internal/dao"
	"one2.3/internal/model/do"
	"one2.3/internal/model/entity"
	"one2.3/internal/service"
)

type sPost struct {
}

func init() {
	service.RegisterPost(New())
}
func New() *sPost {
	return &sPost{}
}

func (s *sPost) List(ctx context.Context, req *sys.PostSearchReq) (res *sys.PostSearchRes, err error) {
	res = new(sys.PostSearchRes)

	m := dao.SysPost.Ctx(ctx)
	if req != nil {
		if req.PostCode != "" {
			m = m.WhereLike(dao.SysPost.Columns().PostCode, "%"+req.PostCode+"%")
		}
		if req.PostName != "" {
			m = m.WhereLike(dao.SysPost.Columns().PostName, "%"+req.PostName+"%")
		}
		if req.Status != "" {
			m = m.Where(dao.SysPost.Columns().Status, gconv.Uint(req.Status))
		}
	}
	res.Total, err = m.Count()
	if err != nil {
		return
	}

	res.Message = consts.OK
	res.Code = 0
	err = m.Page(req.PageNum, req.PageSize).Order("post_sort asc,post_id asc").Scan(&res.PostList)
	if err != nil {
		return
	}
	return
}

func (s *sPost) Add(ctx context.Context, req *sys.PostAddReq) (err error) {

	_, err = dao.SysPost.Ctx(ctx).Insert(do.SysPost{
		PostCode:  req.PostCode,
		PostName:  req.PostName,
		PostSort:  req.PostSort,
		Status:    req.Status,
		Remark:    req.Remark,
		CreatedBy: service.Context().Get(ctx).User.Id,
	})
	service.Cache().Remove(ctx, "sysPost")
	return
}

func (s *sPost) Edit(ctx context.Context, req *sys.PostEditReq) (err error) {

	_, err = dao.SysPost.Ctx(ctx).WherePri(req.PostId).Update(do.SysPost{
		PostCode:  req.PostCode,
		PostName:  req.PostName,
		PostSort:  req.PostSort,
		Status:    req.Status,
		Remark:    req.Remark,
		UpdatedBy: service.Context().Get(ctx).User.Id,
	})
	if err == nil {
		service.Cache().Remove(ctx, "sysPost")
	}
	return
}

func (s *sPost) Delete(ctx context.Context, ids []int) (err error) {

	_, err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().PostId+" in(?)", ids).Delete()
	if err == nil {
		service.Cache().Remove(ctx, "sysPost")
	}
	return
}

// GetPost 获取正常状态的岗位
func (s *sPost) GetPost(ctx context.Context) (list []*entity.SysPost, err error) {
	var cacheList interface{}
	cacheList, err = service.Cache().GetOrSetFuncLock(ctx, "sysPost", func(ctx context.Context) (value interface{}, err error) {
		err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().Status, 1).
			Order(dao.SysPost.Columns().PostSort + " ASC, " + dao.SysPost.Columns().PostId + " ASC ").Scan(&list)
		if err != nil {
			return nil, err
		}
		value = list
		return
	}, 0)
	if cacheList != nil {
		err = gconv.Struct(cacheList, &list)
		if err != nil {
			return nil, err
		}

	}

	return
}
