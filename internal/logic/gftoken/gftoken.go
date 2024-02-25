package gftoken

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"one2.3/api/v1/sys"
	"one2.3/internal/dao"
	"one2.3/internal/model"
	"one2.3/internal/model/entity"
	"one2.3/internal/service"
	"one2.3/utility/gtoken"
	"one2.3/utility/response"
)

type sGftoken struct {
	*gtoken.GfToken
}

func New() *sGftoken {
	ctx := gctx.New()
	cacheK := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	return &sGftoken{GfToken: &gtoken.GfToken{
		AuthAfterFunc: AuthAfter,
		//AuthExcludePaths: g.SliceStr{"/sys/user/login", "/sys/user/logout"},
		CacheMode: 2,
		CacheKey:  cacheK,
	},
	}
}

func init() {
	service.RegisterGftoken(New())
}
func (s *sGftoken) GetTokenData(ctx context.Context) gtoken.Resp {
	r := g.RequestFromCtx(ctx)
	return s.GfToken.GetTokenData(r)

}
func (s *sGftoken) Middleware(ctx context.Context, group *ghttp.RouterGroup) error {
	err := s.GfToken.Middleware(ctx, group)
	return err
}
func (s *sGftoken) GenToken(ctx context.Context, userKey string, data interface{}) gtoken.Resp {
	return s.GfToken.GenToken(ctx, userKey, data)
}
func (s *sGftoken) DownUser(ctx context.Context, UserName string) bool {

	return s.GfToken.DownUser(ctx, UserName)
}
func (s *sGftoken) Login(ctx context.Context, in *model.UserLoginInput) (res *sys.UserLoginRes, err error) {
	capRes := service.Captcha().Verify(in.IdKey, in.Captcha)

	if !capRes {
		response.Fail(ctx, 0, "请输入正确的验证码")
	}

	user, err := s.CheckUserPassword(ctx, model.Check{Username: in.Username, Password: in.Password})
	if err != nil {
		return nil, err
	}

	user.UserPassword = ""
	respToken := s.GenToken(ctx, user.UserName, user)
	if !respToken.Success() {
		err = gerror.New("创建token失败,redis正常否？")
		return nil, err
	}
	user.UserSalt = ""
	res = &sys.UserLoginRes{
		Mobile:       user.Mobile,
		UserNickname: user.UserNickname,
		Avatar:       user.Avatar,
		DeptId:       user.DeptId,
		PostId:       user.PostId,
		Token:        respToken.GetString("token"),
	}
	return res, nil
}
func (s *sGftoken) GetMemberToken(ctx context.Context, memberInfo *entity.Member) (token string, err error) {
	respToken := s.GenToken(ctx, memberInfo.Openid, memberInfo)
	if !respToken.Success() {
		err = gerror.New("创建token失败,redis正常否？")
		return "", err
	}
	return respToken.GetString("token"), err
}
func (s *sGftoken) Logout(ctx context.Context) bool {

	return s.GfToken.DownUser(ctx, service.Context().Get(ctx).User.UserName)

}
func (s *sGftoken) GetOut(ctx context.Context, UserName string) bool {

	return s.GfToken.DownUser(ctx, UserName)
}

/*
	func DoLogin(r *ghttp.Request) (string, interface{}) {
		var (
			ctx = r.GetCtx()
			in  = model.UserLoginInput{}
		)
		if err := r.Parse(&in); err != nil {
			response.Fail(ctx, 0, err.Error())
		}

		capRes := service.Captcha().Verify(in.IdKey, in.Captcha)

		if !capRes {
			response.Fail(ctx, 0, "请输入正确的验证码")
		}

		user, err := CheckUserPassword(ctx, model.Check{Username: in.Username, Password: in.Password})
		if err != nil {
			response.Fail(ctx, 0, err.Error())
		}

		user.UserPassword = ""

		// 登录日志
		//model.UpdateTime = library.GetNow()
		//model.UpdateId = model.Id
		//log.SaveLog(model, constants.LOGIN)

		return user.UserName, user
	}
*/
func AuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		r.Middleware.Next()
	} else {
		var params map[string]interface{}
		//if r.Method == http.MethodGet {
		//	params = r.GetMap()
		//} else if r.Method == http.MethodPost {
		//	params = r.GetMap()
		//} else {
		//	r.Response.Writeln("错误的请求方式")
		//	return
		//}
		params = r.GetMap()

		no := gconv.String(gtime.TimestampMilli())

		g.Log().Warning(r.Context(), fmt.Sprintf("[AUTH_%s][url:%s][params:%s][data:%s]",
			no, r.URL.Path, params, respData.Json()))
		response.JsonExit(r, 1001, "鉴权失败，重新登录")
	}
}
func (s *sGftoken) CheckUserPassword(ctx context.Context, in model.Check) (user *model.CtxUser, err error) {
	user = &model.CtxUser{}
	err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().UserName, in.Username).Scan(&user)
	if user.Id <= 0 {
		err = gerror.New("用户名或密码错误")
		return
	}
	if err != nil {
		return
	}
	if gmd5.MustEncryptString(gmd5.MustEncryptString(in.Password)+gmd5.MustEncryptString(user.UserSalt)) != user.UserPassword {
		err = gerror.New("密码或用户名错误")
	}
	if user.UserStatus != 1 {
		err = gerror.New("状态异常")
	}
	return
}
func (s *sGftoken) GetUser(ctx context.Context) (user *model.CtxUser, err error) {
	var resp gtoken.Resp

	resp = s.GetTokenData(ctx)
	if !resp.Success() {
		return nil, gerror.New("获取用户信息错误")
	}
	err = gjson.DecodeTo(resp.GetString("data"), &user)
	if err != nil {
		return
	}

	return
}
func (s *sGftoken) GetMember(ctx context.Context) (member *entity.Member, err error) {
	var resp gtoken.Resp

	resp = s.GetTokenData(ctx)
	if !resp.Success() {
		return nil, gerror.New("获取用户信息错误")
	}
	err = gjson.DecodeTo(resp.GetString("data"), &member)
	if err != nil {
		return
	}

	return
}
