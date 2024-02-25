package home

import (
	"context"
	"one2.3/api/v1/home"
	"one2.3/internal/dao"
	"one2.3/internal/model/do"
	"one2.3/internal/service"
)

var (
	Member = cMember{}
)

type cMember struct{}

func (c *cMember) GetMemberInfo(ctx context.Context, req *home.MemberInfoReq) (res *home.MemberInfoRes, err error) {

	mem := service.Context().MGet(ctx)
	res = &home.MemberInfoRes{
		Id:         mem.Id,
		Realname:   mem.Realname,
		Nickname:   mem.Nickname,
		Idcard:     mem.Idcard,
		Group:      mem.Group,
		Bigclass:   mem.Bigclass,
		Smallclass: mem.Smallclass,
		Mobile:     mem.Mobile,
		WebAuth:    mem.WebAuth,
		Avatar:     mem.Avatar,
		Sex:        mem.Sex,
	}
	return
}
func (c *cMember) UpdateMemberInfo(ctx context.Context, req *home.UpdateMemberInfoReq) (res *home.UpdateMemberInfoRes, err error) {
	_, err = dao.Member.Ctx(ctx).OmitEmpty().WherePri(service.Context().MGet(ctx).Id).Update(do.Member{
		Id:         req.Id,
		Realname:   req.Realname,
		Nickname:   req.Nickname,
		Idcard:     req.Idcard,
		Group:      req.Group,
		Bigclass:   req.Bigclass,
		Smallclass: req.Smallclass,
		Mobile:     req.Mobile,
		WebAuth:    req.WebAuth,
		Avatar:     req.Avatar,
		Sex:        req.Sex,
	})
	return
}
