package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WechatLogin struct {
	g.Meta `path:"/wechat/login" tags:"微信" method:"get" summary:"登录"`
	//UserName string `p:"userName"      dc:"用户名"`

}
