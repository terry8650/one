package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"one2.3/internal/model/entity"
)

type UserLoginInput struct {
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
	Captcha  string `json:"captcha"  v:"required#请输入验证码"`
	IdKey    string `p:"idkey" v:"required#验证码KEY不能为空"`
}
type Check struct {
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

// SysUserAllRes 带有部门、角色、岗位信息的用户数据
type SysUserAllRes struct {
	gmeta.Meta `orm:"table:sys_user"`
	*entity.SysUser
	RoleInfo []*SysUserRoleInfoRes `json:"roleInfo" `
	Post     []*SysUserPostInfoRes `json:"post"`
}

type SysUserRoleInfoRes struct {
	gmeta.Meta `orm:"table:sys_role"`
	Id         uint   `json:"roleId"`
	Name       string `json:"name"`
}

type SysUserPostInfoRes struct {
	gmeta.Meta `orm:"table:sys_post"`
	PostId     int64  `json:"postId"`
	PostName   string `json:"postName"`
}
