package home

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"one2.3/api/v1/home"
	"one2.3/internal/consts"
	"one2.3/internal/dao"
	"one2.3/internal/model/do"
	"one2.3/internal/service"
	"one2.3/utility/response"
	"sync"
)

type cChildren struct{ lock *sync.Mutex }

var (
	Children = cChildren{lock: &sync.Mutex{}}
)

func (c *cChildren) GetStudent(ctx context.Context, req *home.ChildrenGetListReq) (res *home.ChildrenGetListRes, err error) {
	res = &home.ChildrenGetListRes{}
	err = dao.Student.Ctx(ctx).Where(dao.Student.Columns().Cid, req.Cid).Scan(&res.Stu)

	return
}
func (c *cChildren) CheckInClass(ctx context.Context, req *home.CheckInClassReq) (res *home.CheckInClassRes, err error) {
	res = &home.CheckInClassRes{}
	err = dao.Memclassstu.Ctx(ctx).Where(g.Map{dao.Memclassstu.Columns().Cid: req.Cid, dao.Memclassstu.Columns().Mid: service.Context().MGet(ctx).Id, dao.Memclassstu.Columns().Status: 1}).Scan(&res.MemStuList)
	return
}
func (c *cChildren) AddMem(ctx context.Context, req *home.AddMemReq) (res *home.AddMemRes, err error) {
	var data do.Memclassstu
	gconv.Struct(req, &data)
	data.Mid = service.Context().MGet(ctx).Id
	_, err = dao.Memclassstu.Ctx(ctx).OmitEmpty().Data(data).Insert()
	return
}
func (c *cChildren) CreateZuoye(ctx context.Context, req *home.CreateZuoYeReq) (res *home.CreateZuoYeRes, err error) {
	Children.lock.Lock()
	defer Children.lock.Unlock()
	_, err = dao.Zuoye.Ctx(ctx).OmitEmpty().Data(do.Zuoye{
		Type:          req.Type,
		Cid:           req.Cid,
		Creatormid:    service.Context().MGet(ctx).Id,
		Creatorname:   service.Context().MGet(ctx).Realname,
		Createday:     req.Createday,
		Content:       req.Content,
		Startday:      req.Startday,
		Endday:        req.Endday,
		Lastupdateman: gconv.String(service.Context().MGet(ctx).Realname) + gconv.String(service.Context().MGet(ctx).Id),
		Status:        req.Status,
	}).Insert()
	if err != nil && gerror.Code(err).Code() == 52 {
		response.JsonOri(ctx, g.Map{"code": 0, "msg": "重复创建"})
	}

	return
}
func (c *cChildren) UpdateZuoye(ctx context.Context, req *home.UpdateZuoYeReq) (res *home.UpdateZuoYeRes, err error) {
	var data do.Zuoye
	gconv.Struct(req, &data)
	data.Lastupdateman = gconv.String(service.Context().MGet(ctx).Realname) + gconv.String(service.Context().MGet(ctx).Id)
	_, err = dao.Zuoye.Ctx(ctx).WherePri(req.Id).OmitEmpty().Data(data).Update()
	if err != nil && gerror.Code(err).Code() == 52 {
		response.JsonOri(ctx, g.Map{"code": 0, "msg": "重复创建"})
	}
	return
}
func (c *cChildren) ZuoYeDone(ctx context.Context, req *home.ZuoYeDoneReq) (res *home.ZuoYeDoneRes, err error) {
	var data do.Zuoyedone
	data.Mid = service.Context().MGet(ctx).Id
	gconv.Struct(req, &data)
	_, err = dao.Zuoyedone.Ctx(ctx).WherePri(req.Zyid).OmitEmpty().Data(data).Insert()
	return
}

func (c *cChildren) SaveUpload(ctx context.Context, req *home.SaveUploadReq) (res *home.SaveUploadRes, err error) {
	_, err = dao.Zuoye.Ctx(ctx).WherePri(req.Id).Update(do.Zuoye{Files: req.Imgs})
	return
}
func (c *cChildren) ZuoYeList(ctx context.Context, req *home.ZuoYeListReq) (res *home.ZuoYeListRes, err error) {
	res = &home.ZuoYeListRes{}

	m := dao.Zuoye.Ctx(ctx)
	if req != nil {
		if req.Cid != "" {
			m = m.Where(dao.Zuoye.Columns().Cid, req.Cid)
		}
		if req.Type != "" {
			m = m.Where(dao.Zuoye.Columns().Type, req.Type)
		}
		if req.Theday != "" && req.Theday != "all" {
			m = m.Where(g.Map{
				"startday <=": req.Theday,
				"endday >=":   req.Theday,
			})
		}

	}
	res.Total, err = m.Count()
	if err != nil {
		return
	}

	//m = m.LeftJoin("zuoyedone zy", "zuoye.id=zy.zyid").Fields("zuoye.*,zy.mid,zy.stuname,zy.stucall").Where("zy.stuid=? OR zy.stuid is null", req.Stuid)

	res.Message = consts.OK
	res.Code = 0
	err = m.Page(req.PageNum, req.PageSize).Order("id desc").Scan(&res.ZuoYeList)
	arr2 := make([]*uint64, len(res.ZuoYeList))
	for i := range res.ZuoYeList {
		arr2[i] = &res.ZuoYeList[i].Id
	}
	err = dao.Zuoyedone.Ctx(ctx).WhereIn(dao.Zuoyedone.Columns().Zyid, arr2).Where(dao.Zuoyedone.Columns().Stuid, req.Stuid).Scan(&res.ZuoYeDoneList)
	if err == nil {
		response.JsonOri(ctx, res)

	}
	return
}
