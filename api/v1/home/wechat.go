package home

import (
	"github.com/gogf/gf/v2/frame/g"
	"one2.3/internal/model/entity"
)

type WechatLoginReq struct {
	g.Meta `path:"/wechat/login" tags:"微信" method:"get" summary:"登录"`
	Tag    string `p:"tag"      dc:"跳转标记"`
}
type WechatLoginRes struct {
}
type CallBackReq struct {
	g.Meta `path:"/wechat/callback" tags:"微信" method:"get" summary:"userinfo"`
	Code   string `p:"code"`
	Tag    string `p:"tag"`
}
type CallBackRes struct {
	Token string `json:"mToken"`
	Info  *entity.Member
}
type JsReq struct {
	g.Meta `path:"/wechat/get_js" tags:"微信" method:"post" summary:"userinfo"`
	Url    string
}
type JsRes struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}
