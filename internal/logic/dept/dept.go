package dept

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"one2.3/api/v1/sys"
	"one2.3/internal/consts"
	"one2.3/internal/dao"
	"one2.3/internal/model/do"
	"one2.3/internal/service"
)

type sDept struct {
}

func init() {
	service.RegisterDept(New())
}
func New() *sDept {
	return &sDept{}
}
func (s *sDept) List(ctx context.Context) (res *sys.DeptListRes, err error) {
	res = &sys.DeptListRes{}
	var cacheList interface{}
	cacheList, err = service.Cache().GetOrSetFuncLock(ctx, "sysDept", func(ctx context.Context) (value interface{}, err error) {
		err = dao.SysDept.Ctx(ctx).Scan(&res.DeptList)
		if err != nil {
			return nil, err
		}
		value = res.DeptList
		return
	}, 0)
	if err != nil {
		return
	}
	if cacheList != nil {
		err = gconv.Struct(cacheList, &res.DeptList)
		if err != nil {
			return nil, err
		}

	}
	res.Message = consts.OK
	res.Code = 0

	return
}
func (s *sDept) Add(ctx context.Context, req *sys.DeptAddReq) (err error) {
	_, err = dao.SysDept.Ctx(ctx).Insert(do.SysDept{
		Pid:       req.Pid,
		Name:      req.Name,
		Sort:      req.Sort,
		Tel:       req.Tel,
		Email:     req.Email,
		Status:    req.Status,
		CreatedBy: service.Context().Get(ctx).User.Id,
	})
	if err == nil {
		service.Cache().Remove(ctx, "sysDept")
	}

	return
}
func (s *sDept) Update(ctx context.Context, req *sys.DeptUpdateReq) (err error) {
	_, err = dao.SysDept.Ctx(ctx).WherePri(req.Id).Update(do.SysDept{
		Pid:       req.Pid,
		Name:      req.Name,
		Sort:      req.Sort,
		Tel:       req.Tel,
		Email:     req.Email,
		Status:    req.Status,
		UpdatedBy: service.Context().Get(ctx).User.Id,
	})
	if err == nil {
		service.Cache().Remove(ctx, "sysDept")
	}
	return
}
func (s *sDept) Delete(ctx context.Context, id int64) (err error) {
	if s.existSon(ctx, id) {
		err = gerror.New("先删除子节点")
		return
	}
	_, err = dao.SysDept.Ctx(ctx).Where(do.SysDept{Id: id}).Delete()
	if err == nil {
		service.Cache().Remove(ctx, "sysDept")
	}

	return

}
func (s *sDept) existSon(ctx context.Context, id int64) bool {

	d, err := dao.SysDept.Ctx(ctx).Where(do.SysDept{Pid: id}).One()

	if err != nil {
		g.Log().Error(ctx, err)
		return false
	}

	return !d.IsEmpty()
}
