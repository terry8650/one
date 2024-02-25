package home

import (
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/api/v1/common"
	"one2.3/internal/model/entity"
)

type ChildrenGetListReq struct {
	g.Meta `path:"/children/list" tags:"children" method:"get" summary:"class students"`
	Cid    string `p:"cid"      dc:"class-id"`
}
type ChildrenGetListRes struct {
	Stu []*entity.Student `json:"list"`
}
type CreateClassReq struct {
	g.Meta `path:"/children/addclass" tags:"children" method:"get" summary:"class students"`
	*entity.Cla
}
type CreateClassRes struct {
}
type CheckInClassReq struct {
	g.Meta `path:"/children/checkinclass" tags:"children" method:"get" summary:"核查是否本班的"`
	Cid    string `p:"cid"      dc:"class-id"`
}
type CheckInClassRes struct {
	MemStuList []*entity.Memclassstu `json:"list"`
}
type ZuoYeBase struct {
	Type string `p:"type"          description:"作业类型"`
	Cid  string `p:"cid"           description:"班级ID"`

	Createday string `p:"createday"     description:""`
	Content   string `p:"content"       description:""`
	Startday  string `p:"startday"      description:""`
	Endday    string `p:"endday"        description:""`

	Status string `p:"status" description:"状态1显示2隐藏"`
}
type AddMemReq struct {
	g.Meta `path:"/children/addmem" tags:"children" method:"post" summary:"添加人员"`

	//Mid       uint64      `json:"mid"       description:""`
	Mname string `json:"mname"     description:""`
	Cid   string `json:"cid"       description:"班级ID"`
	//Cname     string      `json:"cname"     description:"班级名称"`
	Stuid   int64  `json:"stuid"     description:""`
	Stuname string `json:"stuname"   description:"多个孩子用多条数据"`
	Stucall string `json:"stucall"   description:"和孩子的关系"`
	Mtel    string `json:"mtel"      description:""`
}
type AddMemRes struct {
}
type CreateZuoYeReq struct {
	g.Meta `path:"/children/addzuoye" tags:"children" method:"post" summary:"添加作业"`
	ZuoYeBase
}
type CreateZuoYeRes struct {
}
type UpdateZuoYeReq struct {
	g.Meta `path:"/children/updatezuoye" tags:"children" method:"put" summary:"修改作业"`
	ZuoYeBase
	Id uint64 `p:"id"     v:"required#id不能为空"       description:""`
}
type UpdateZuoYeRes struct {
}
type ZuoYeDoneBase struct {
	Zyid uint64 `json:"zyid"            description:""`

	Mid         int64  `json:"mid"         description:""`
	Stuid       int64  `json:"stuid"       description:""`
	Stuname     string `json:"stuname"     description:""`
	Stucall     string `json:"stucall"     description:""`
	Donecontent string `json:"donecontent" description:""`
	//DoneStatus      int         `json:"donestatus"      description:"1完成"`
	Zyzl string `json:"zyzl"        description:"作业质量"`
}
type ZuoYeDoneReq struct {
	g.Meta `path:"/children/donezuoye" tags:"children" method:"put" summary:"完成作业"`

	Mid         int64  `p:"mid"         description:""`
	Stuid       int64  `p:"stuid"       description:""`
	Stuname     string `p:"stuname"     description:""`
	Stucall     string `p:"stucall"     description:""`
	Donecontent string `p:"donecontent" description:""`
	Zyid        uint64 `p:"zyid"     v:"required#id不能为空"       description:""`
}
type ZuoYeDoneRes struct {
}
type UploadZyReq struct {
	g.Meta `path:"/children/uploadzuoye" tags:"children" method:"post" summary:"上传作业"`
	Sid    string `p:"sid"      dc:"serid"`
}

type UploadZyRes struct {
	Url string `json:"url"  `
}
type SaveUploadReq struct {
	g.Meta `path:"/children/saveupload" tags:"children" method:"put" summary:"save上传"`
	Imgs   string `p:"imgs"      `
	Id     string `p:"id"   `
}
type SaveUploadRes struct {
}
type ZuoYeListReq struct {
	g.Meta `path:"/children/zuoyelist" tags:"children" method:"get" summary:"作业list"`
	Cid    string `p:"cid"      v:"required#兄弟哪个班的？"    description:"班级"`
	Theday string `p:"theday"          description:"作业日期"`
	Type   string `p:"type"          description:"作业类型"`
	Stuid  int64  `p:"stuid"       description:""`
	common.PageReq
}
type ZuoYeListRes struct {
	ZuoYeList     []*entity.Zuoye  `json:"data"`
	ZuoYeDoneList []*ZuoYeDoneBase `json:"donelist"`
	common.ListRes
}
