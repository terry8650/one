package home

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MemberInfoReq struct {
	g.Meta `path:"/member/info" tags:"member信息" method:"get" summary:"userinfo"`
}
type MemberInfoRes struct {
	Id         uint64 `json:"id"          description:""`
	Realname   string `json:"realname"    description:""`
	Nickname   string `json:"nickname"    description:""`
	Idcard     string `json:"idcard"      description:""`
	Group      string `json:"group"       description:""`
	Bigclass   string `json:"bigclass"    description:""`
	Smallclass string `json:"smallclass"  description:""`
	Mobile     string `json:"mobile"      description:""`
	WebAuth    string `json:"webAuth"     description:""`
	Avatar     string `json:"avatar"      description:""`
	Sex        int    `json:"sex"         description:"0woman1man"`
}
type UpdateMemberInfoReq struct {
	g.Meta     `path:"/member/update" tags:"member信息" method:"put" summary:"userinfo"`
	Id         uint64 `p:"id"          description:""`
	Realname   string `p:"realname"    description:""`
	Nickname   string `p:"nickname"    description:""`
	Idcard     string `p:"idcard"      description:""`
	Group      string `p:"group"       description:""`
	Bigclass   string `p:"bigclass"    description:""`
	Smallclass string `p:"smallclass"  description:""`
	Mobile     string `p:"mobile"      description:""`
	WebAuth    string `p:"webAuth"     description:""`
	Avatar     string `p:"avatar"      description:""`
	Sex        int    `p:"sex"         description:"0woman1man"`
}
type UpdateMemberInfoRes struct {
}
