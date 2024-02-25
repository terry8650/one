package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/api/v1/common"
	"one2.3/internal/model/entity"
)

// UserSearchReq 用户搜索请求参数
type UserSearchReq struct {
	g.Meta   `path:"/user/list" tags:"用户管理" method:"get" summary:"用户列表"`
	UserName string `p:"userName"      dc:"用户名"`
	DeptId   string `p:"deptId"` //部门id
	Mobile   string `p:"mobile"`
	Status   string `p:"status"`
	PostId   string `p:"postId"`
	common.PageReq
}
type UserBase struct {
	UserName     string `json:"userName"      dc:"用户名" v:"required#username不能为空"`
	Mobile       string `json:"mobile"        dc:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname string `json:"userNickname"  dc:"用户昵称"`
	Birthday     int    `json:"birthday"      dc:"生日"`
	UserStatus   uint   `json:"userStatus"    dc:"用户状态;0:禁用,1:正常,2:未验证"`
	UserEmail    string `json:"userEmail"     dc:"用户登录邮箱"`
	Sex          int    `json:"sex"           dc:"性别;0:保密,1:男,2:女"`
	Avatar       string `json:"avatar"        dc:"用户头像"`
	DeptId       uint64 `json:"deptId"        dc:"部门id"`
	Remark       string `json:"remark"        dc:"备注"`
	UserSalt     string `json:"salt"`
	PostId       uint64 `json:"postId"`
	Address      string `json:"address"       dc:"联系地址"`
	Describe     string `json:"describe"      dc:"描述信息"`
	LastLoginIp  string `json:"lastLoginIp"   dc:"最后登录ip"`
}
type UserListBase struct {
	UserBase
	Id uint64 `json:"id" v:"required#Id不能为空"`
}
type UserSearchRes struct {
	g.Meta   `mime:"application/json"`
	UserList []*UserListBase `json:"data"`
	common.ListRes
}
type UserGetDeptPostReq struct {
	g.Meta `path:"/user/get_dept_post_role" tags:"用户管理" method:"get" summary:"获取用户部门职位"`
}
type UserGetDeptPostRes struct {
	DeptList []*entity.SysDept `json:"dept"`
	PostList []*entity.SysPost `json:"post"`
	RoleList []*entity.SysRole `json:"role"`
}
type UserGetRoleReq struct {
	g.Meta `path:"/user/get_role_ids" tags:"用户管理" method:"get" summary:"获取用户角色ids"`
	Id     uint64 `json:"id" v:"required#id不能为空"`
}
type UserGetRoleRes struct {
	RoleIds []uint `json:"roleIds"`
}
type UserAddReq struct {
	g.Meta `path:"/user/add" tags:"用户管理" method:"post" summary:"添加用户"`
	UserBase
	RoleIds []uint `p:"roleIds" v:"required#角色id不能为空"`
}
type UserAddRes struct {
}
type UserUpdateReq struct {
	g.Meta `path:"/user/update" tags:"用户管理" method:"put" summary:"更新用户"`
	UserListBase
	RoleIds []uint `p:"roleIds" v:"required#角色id不能为空"`
}
type UserUpdateRes struct {
}
type UserDeleteReq struct {
	g.Meta `path:"/user/delete" tags:"用户管理" method:"delete" summary:"删除用户"`
	Id     uint64 `json:"id" v:"required#userid不能为空"`
}
type UserDeleteRes struct {
}

type UserChangePwdReq struct {
	g.Meta `path:"/user/change_pwd" tags:"用户管理" method:"put" summary:"修改密码"`

	Password string `p:"password" v:"required#密码不能为空"`
	UserName string `p:"username" v:"required#用户名不能为空"`
	Salt     string `p:"salt" v:"required#salt不能为空"`
}
type UserChangePwdRes struct {
}
