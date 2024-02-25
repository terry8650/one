// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta    `orm:"table:sys_menu, do:true"`
	Id        interface{} //
	Title     interface{} // 菜单名称
	Name      interface{} //
	Jump      interface{} //
	Url       interface{} //
	Icon      interface{} //
	Remark    interface{} //
	Pid       interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}