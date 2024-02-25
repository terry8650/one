// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"one2.3/api/v1/sys"
	"one2.3/internal/model"
	"one2.3/internal/model/entity"
	"one2.3/utility/gtoken"
)

type (
	IGftoken interface {
		GetTokenData(ctx context.Context) gtoken.Resp
		Middleware(ctx context.Context, group *ghttp.RouterGroup) error
		GenToken(ctx context.Context, userKey string, data interface{}) gtoken.Resp
		DownUser(ctx context.Context, UserName string) bool
		Login(ctx context.Context, in *model.UserLoginInput) (res *sys.UserLoginRes, err error)
		GetMemberToken(ctx context.Context, memberInfo *entity.Member) (token string, err error)
		Logout(ctx context.Context) bool
		GetOut(ctx context.Context, UserName string) bool
		CheckUserPassword(ctx context.Context, in model.Check) (user *model.CtxUser, err error)
		GetUser(ctx context.Context) (user *model.CtxUser, err error)
		GetMember(ctx context.Context) (member *entity.Member, err error)
	}
)

var (
	localGftoken IGftoken
)

func Gftoken() IGftoken {
	if localGftoken == nil {
		panic("implement not found for interface IGftoken, forgot register?")
	}
	return localGftoken
}

func RegisterGftoken(i IGftoken) {
	localGftoken = i
}
