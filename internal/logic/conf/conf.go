package conf

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"one2.3/api/v1/sys"
	"one2.3/internal/dao"
	"one2.3/internal/model/do"
	"one2.3/internal/service"
)

type sConf struct {
}

func init() {
	service.RegisterConf(New())
}
func New() *sConf {
	return &sConf{}
}
func (s *sConf) List(ctx context.Context) (res *sys.ConfListRes, err error) {
	res = &sys.ConfListRes{}
	var cacheList interface{}
	cacheList, err = service.Cache().GetOrSetFuncLock(ctx, "sysConf", func(ctx context.Context) (value interface{}, err error) {
		err = dao.SysConfig.Ctx(ctx).Scan(&res.ConfList)
		if err != nil {
			return nil, err
		}
		value = res.ConfList
		return
	}, 0)
	if err != nil {
		return
	}
	if cacheList != nil {
		err = gconv.Struct(cacheList, &res.ConfList)
		if err != nil {
			return nil, err
		}
	}

	return
}
func (s *sConf) GetVal(ctx context.Context, req *sys.ConfGetReq) (res *sys.ConfGetRes, err error) {
	res = &sys.ConfGetRes{}
	var d interface{}
	d, err = dao.SysConfig.Ctx(ctx).Fields(dao.SysConfig.Columns().Val).Where(do.SysConfig{Name: req.Name}).Value()
	res.Val = gconv.String(d)
	return

}
func (s *sConf) Add(ctx context.Context, req *sys.ConfAddReq) (err error) {
	_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{
		Name:   req.Name,
		Val:    req.Val,
		Remark: req.Remark,
	})
	if err == nil {
		service.Cache().Remove(ctx, "sysConf")
	}

	return
}
func (s *sConf) Update(ctx context.Context, req *sys.ConfUpdateReq) (err error) {
	_, err = dao.SysConfig.Ctx(ctx).WherePri(req.Id).Update(do.SysConfig{

		Id:     req.Id,
		Name:   req.Name,
		Val:    req.Val,
		Remark: req.Remark,
	})
	if err == nil {
		service.Cache().Remove(ctx, "sysConf")
	}
	return
}
func (s *sConf) Delete(ctx context.Context, ids []uint) (err error) {
	_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().Id+" in(?)", ids).Delete()
	if err == nil {
		service.Cache().Remove(ctx, "sysConf")
	}
	return
}
