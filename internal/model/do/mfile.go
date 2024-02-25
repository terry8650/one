// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Mfile is the golang structure of table mfile for DAO operations like Where/Data.
type Mfile struct {
	g.Meta    `orm:"table:mfile, do:true"`
	Id        interface{} // 自增ID
	Name      interface{} // 文件名称
	Src       interface{} // 本地文件存储路径
	Url       interface{} // URL地址，可能为空
	MemberId  interface{} // 操作用户
	CreatedAt *gtime.Time //
}