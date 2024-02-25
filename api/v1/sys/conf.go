package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/internal/model/entity"
)

type ConfListReq struct {
	g.Meta `path:"/conf/list" tags:"系统参数管理" method:"get" summary:"参数列表"`
}
type ConfListRes struct {
	g.Meta   `mime:"application/json"`
	ConfList []*entity.SysConfig `json:"data"`
}
type ConfGetReq struct {
	g.Meta `path:"/conf/get_val" tags:"系统参数管理" method:"get" summary:"获取参数"`
	Name   string `p:"name" v:"required#参数名不能为空"`
}
type ConfGetRes struct {
	g.Meta `mime:"application/json"`
	Val    string `json:"val"`
}
type ConfAddReq struct {
	g.Meta `path:"/conf/add" tags:"系统参数管理" method:"post" summary:"添加参数"`
	Name   string `json:"name"      v:"required#参数名称不能为空"`
	Val    string `json:"val"       v:"required#参数值不能为空"`
	Remark string `json:"remark"`
}
type ConfAddRes struct{}
type ConfUpdateReq struct {
	g.Meta `path:"/conf/update" tags:"系统参数管理" method:"put" summary:"更新参数"`
	Id     uint   `json:"id"        v:"required#id不能为空"`
	Name   string `json:"name"      v:"required#参数名称不能为空"`
	Val    string `json:"val"       v:"required#参数值不能为空"`
	Remark string `json:"remark"`
}
type ConfUpdateRes struct{}
type ConfDeleteReq struct {
	g.Meta `path:"/conf/delete" tags:"系统参数管理" method:"delete" summary:"删除参数"`
	Ids    []uint `p:"ids" v:"required#ids不能为空"`
}
type ConfDeleteRes struct{}
