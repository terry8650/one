package model

import "github.com/gogf/gf/v2/frame/g"

// Context 请求上下文结构
type Context struct {
	User *CtxUser // 上下文用户信息
	Data g.Map    // 自定KV变量，业务模块根据需要设置，不固定
}
type CtxUser struct {
	Id           uint64 `json:"id"            description:""`
	UserName     string `json:"userName"      description:"用户名"`
	Mobile       string `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname string `json:"userNickname"  description:"用户昵称"`
	UserSalt     string `json:"userSalt"      description:"加密盐"`
	UserPassword string `json:"userPassword"  description:"登录密码;cmf_password加密"`
	UserStatus   uint   `json:"userStatus"    description:"用户状态;0:禁用,1:正常,2:未验证"`
	Avatar       string `json:"avatar"        description:"用户头像"`
	DeptId       uint64 `json:"deptId"        description:"部门id"`
	PostId       uint64 `json:"postId"        description:"职位id"`
}
