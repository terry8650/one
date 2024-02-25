package sys

import "github.com/gogf/gf/v2/frame/g"

type PersonalChangePwdReq struct {
	g.Meta      `path:"/personal/change_pwd" tags:"personal" method:"put" summary:"修改密码"`
	OldPassWord string `p:"oldPassword" v:"required#原密码不能为空"`
	Password    string `p:"password" v:"required#密码不能为空"`
	RePassword  string `p:"repassword" v:"required#密码不能为空|eq:Password"`
}
type PersonalChangePwdRes struct {
}
type PersonalLogoutReq struct {
	g.Meta `path:"/personal/logout" tags:"用户管理" method:"post" summary:"用户登出"`
}
type PersonalLogoutRes struct {
}
