package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/api/v1/common"
	"one2.3/internal/model/entity"
)

type MenuBase struct {
	Title     string `p:"title" v:"required#菜单名称不能为空"`
	Name      string `p:"name"`
	Jump      string
	Icon      string
	Remark    string
	Pid       string `d:"0" v:"required#父菜单不能为空"`
	Type      uint   `json:"type"      description:"类型1菜单 2接口"`
	SmallAuth string `json:"smallAuth" description:"接口细化权限说明，用来写入casbin第二个权限"`
	Condition string `json:"condition" description:"条件：nocheck就不用检测权限"`
}
type MenuAuthList struct {
	*entity.SysAuthRule
	MyAuth string `json:"myAuth"`
}
type MenuAddReq struct {
	g.Meta `path:"/menu/add" tags:"菜单管理" method:"post" summary:"添加菜单"`
	MenuBase
}
type MenuAddRes struct{}
type MenuEditReq struct {
	g.Meta `path:"/menu/update" tags:"菜单管理" method:"put" summary:"更新菜单"`
	Id     uint `p:"id" v:"required#id必须"`
	MenuBase
}
type MenuEditRes struct{}
type MenuSearchReq struct {
	g.Meta `path:"/menu/list" tags:"菜单管理" method:"get" summary:"菜单列表"`
	Title  string
	Name   string
}
type MenuSearchRes struct {
	common.ListRes
	MenuList []*entity.SysAuthRule `json:"data"`
}
type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete" tags:"菜单管理" method:"delete" summary:"删除菜单"`
	Id     uint `p:"id" v:"required#id必须"`
}
type MenuDeleteRes struct{}
type MenuAuthListReq struct {
	g.Meta `path:"/menu/auth_list" tags:"菜单管理" method:"get" summary:"获取权限菜单"`
}
type MenuAuthListRes struct {
	AuthList []*MenuAuthList
}
