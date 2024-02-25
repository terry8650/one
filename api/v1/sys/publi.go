package sys

import "github.com/gogf/gf/v2/frame/g"

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"用户管理" method:"post" summary:"用户登录"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
	Captcha  string `json:"captcha"  v:"required#请输入验证码"`
	IdKey    string `p:"idkey" v:"required#验证码KEY不能为空"`
}
type UserLoginRes struct {
	Mobile       string `json:"mobile"        dc:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname string `json:"userNickname"  dc:"用户昵称"`
	Avatar       string `json:"avatar"        dc:"用户头像"`
	DeptId       uint64 `json:"deptId"        dc:"部门id"`
	PostId       uint64 `json:"postId"`
	Token        string `json:"token"`
}
type TestReq struct {
	g.Meta `path:"/test" tags:"test" method:"post" summary:"test"`
	//Username string `p:"username" v:"required#用户名不能为空"`
	//Password string `p:"password" v:"required#密码不能为空"`
	//Captcha  string `json:"captcha"  v:"required#请输入验证码"`
	//IdKey    string `p:"idkey" v:"required#验证码KEY不能为空"`
}
type TestRes struct {
}
