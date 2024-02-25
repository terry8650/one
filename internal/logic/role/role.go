package role

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

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}
func New() *sRole {
	return &sRole{}
}
func (s *sRole) List(ctx context.Context, req *sys.RoleListReq) (res *sys.RoleListRes, err error) {
	res = new(sys.RoleListRes)

	m := dao.SysRole.Ctx(ctx)
	if req != nil {
		if req.Status != "" {
			m = m.Where(dao.SysRole.Columns().Status, req.Status)
		}
		if req.Name != "" {
			m = m.WhereLike(dao.SysRole.Columns().Name, "%"+req.Name+"%")
		}

	}
	res.Total, err = m.Count()
	if err != nil {
		return
	}

	res.Message = consts.OK
	res.Code = 0
	err = m.Page(req.PageNum, req.PageSize).Order("sort asc,id asc").Scan(&res.RoleList)
	if err != nil {
		return
	}
	return
}
func (s *sRole) Add(ctx context.Context, req *sys.RoleAddReq) (err error) {
	_, err = dao.SysRole.Ctx(ctx).Insert(do.SysRole{
		Status: req.Status,
		Sort:   req.Sort,
		Name:   req.Name,
		Remark: req.Remark,
	})
	if err == nil {
		service.Cache().Remove(ctx, "sysRole")
	}
	return
}
func (s *sRole) Update(ctx context.Context, req *sys.RoleUpdateReq) (err error) {
	_, err = dao.SysRole.Ctx(ctx).WherePri(req.Id).Update(do.SysRole{

		Status: req.Status,
		Sort:   req.Sort,
		Name:   req.Name,
		Remark: req.Remark,
	})
	if err == nil {
		service.Cache().Remove(ctx, "sysRole")
	}
	return

}
func (s *sRole) Delete(ctx context.Context, ids []int) (err error) {
	_, err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id+" in(?)", ids).Delete()
	if err == nil {
		service.Cache().Remove(ctx, "sysRole")
	}
	return
}
func (s *sRole) GetRole(ctx context.Context) (list []*entity.SysRole, err error) {
	var cacheList interface{}
	cacheList, err = service.Cache().GetOrSetFuncLock(ctx, "sysRole", func(ctx context.Context) (value interface{}, err error) {
		err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Status, 1).Scan(&list)
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
