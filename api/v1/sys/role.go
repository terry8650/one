package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/api/v1/common"
	"one2.3/internal/model/entity"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" tags:"角色管理" method:"get" summary:"角色列表"`
	Name   string `json:"name"      description:"角色名称"`
	Status string `p:"status"` //状态
	common.PageReq
}
type RoleListRes struct {
	g.Meta `mime:"application/json"`
	common.ListRes
	RoleList []*entity.SysRole `json:"data"`
}
type RoleAddReq struct {
	g.Meta `path:"/role/add" tags:"角色管理" method:"post" summary:"添加角色"`
	Status uint   `json:"status"    description:"状态;0:禁用;1:正常"`
	Sort   uint   `json:"sort"      description:"排序"`
	Name   string `json:"name"      description:"角色名称"`
	Remark string `json:"remark"    description:"备注"`
}
type RoleAddRes struct {
}
type RoleUpdateReq struct {
	g.Meta `path:"/role/update" tags:"角色管理" method:"put" summary:"修改角色"`
	Id     uint   `json:"id"        description:""`
	Status uint   `json:"status"    description:"状态;0:禁用;1:正常"`
	Sort   uint   `json:"sort"      description:"排序"`
	Name   string `json:"name"      description:"角色名称"`
	Remark string `json:"remark"    description:"备注"`
}
type RoleUpdateRes struct {
}
type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" tags:"角色管理" method:"delete" summary:"删除角色"`
	Ids    []int `p:"ids" v:"required#角色id不能为空"`
}
type RoleDeleteRes struct {
}

type RoleNodesReq struct {
	g.Meta `path:"/role/nodes" tags:"角色管理" method:"get" summary:"获取角色对应节点"`
	RoleId uint `p:"id" v:"required#角色id不能为空"`
}
type RoleNodesRes struct {
	//Role  *entity.SysRole `json:"role"`
	Auth [][]string `json:"auth"`
}
type UpdateOldRulesReq struct {
	g.Meta  `path:"/role/update_old_rule" tags:"角色管理" method:"put" summary:"批量更新角色详细权限"`
	RoleId  uint             `p:"id" v:"required#角色id不能为空"`
	NewAuth map[int][]string `json:"newAuth" v:"required#NewAuth不能为空"`
}
type UpdateOldRulesRes struct {
}
type UpdateSingleRuleReq struct {
	g.Meta  `path:"/role/update_single_rule" tags:"角色管理" method:"put" summary:"更新角色单个详细权限"`
	OldAuth []string `json:"oldAuth" v:"required#OldAuth不能为空"`
	NewAuth []string `json:"newAuth" v:"required#NewAuth不能为空"`
}
type UpdateSingleRuleRes struct {
}
type DelOneRuleReq struct {
	g.Meta `path:"/role/del_one_rule" tags:"角色管理" method:"delete" summary:"删除角色单个权限"`
	Auth   []string `json:"auth" v:"required#auth不能为空"`
}
type DelOneRuleRes struct {
}
type RoleNodesUpdateReq struct {
	g.Meta `path:"/role/save_nodes" tags:"角色管理" method:"post" summary:"更新角色对应节点"`
	RoleId uint   `p:"id" v:"required#角色id不能为空"`
	Nodes  []uint `p:"nodes" v:"required#nodes不能为空"`
}
type RoleNodesUpdateRes struct {
}
type RoleApiNodesReq struct {
	g.Meta `path:"/role/api_nodes" tags:"角色管理" method:"post" summary:"获取角色对应节点"`
	RoleId string `p:"id" v:"required#角色id不能为空"`
}
type RoleApiNodesRes struct {
	Role  *entity.SysRole `json:"role"`
	Nodes []int           `json:"nodes"`
}
