package home

import "github.com/gogf/gf/v2/frame/g"

type HMenuReq struct {
	g.Meta `path:"/menu/list" tags:"微信" method:"get" summary:"userinfo"`
}
type HMenuRes struct {
	AuthList []string
}
