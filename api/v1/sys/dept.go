package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/internal/model/entity"
)

type DeptListReq struct {
	g.Meta `path:"/dept/list" tags:"部门管理" method:"get" summary:"部门列表"`
}

type DeptListRes struct {
	g.Meta  `mime:"application/json"`
	Code    int    `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Message string `json:"msg"`
	//CurrentPage int    `json:"page"`

	DeptList []*entity.SysDept `json:"data"`
}
type DeptBase struct {
	Pid    int64  `json:"pid"       v:"required#父级不能为空"`
	Name   string `json:"name"      v:"required#部门名称不能为空"`
	Sort   int    `json:"sort"`
	Tel    string `json:"tel"  `
	Email  string `json:"email" v:"email#邮箱格式不正确"`
	Status uint   `json:"status" v:"required#状态必须"`
}

type DeptAddReq struct {
	g.Meta `path:"/dept/add" tags:"部门管理" method:"post" summary:"添加部门"`
	DeptBase
}

type DeptAddRes struct {
}

type DeptUpdateReq struct {
	g.Meta `path:"/dept/update" tags:"部门管理" method:"put" summary:"更新部门"`
	Id     int64 `json:"id"        v:"required#id不能为空"`
	DeptBase
}

type DeptUpdateRes struct {
}

type DeptDeleteReq struct {
	g.Meta `path:"/dept/delete" tags:"部门管理" method:"delete" summary:"删除部门"`
	Id     int64 `p:"id" v:"required#id不能为空"`
}

type DeptDeleteRes struct {
}
