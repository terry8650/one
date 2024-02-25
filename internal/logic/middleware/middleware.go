package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"one2.3/internal/model"
	"one2.3/internal/model/entity"
	"one2.3/internal/service"
	"one2.3/utility/response"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}
func (s *sMiddleware) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// you can set options
	//corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()

}
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		//response.JsonExit(r, 0, err.Error())
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			//code = gcode.CodeInternalError
			code = gcode.CodeOK
		}
		//if r.IsAjaxRequest() {
		response.JsonExit(r, code.Code(), err.Error())

		//} else {
		//	//service.View().Render500(r.Context(), model.View{
		//	//	Error: err.Error(),
		//	//})
		//}
	} else {
		response.JsonExit(r, code.Code(), "ok", res)
		//if r.IsAjaxRequest() {
		// 跨域的话，这里是不能判断ajax的
		//	response.JsonExit(r, code.Code(), "", res)
		//} else {
		//	//什么都不做，业务API自行处理模板渲染的成功逻辑。
		//
		//}
	}
}
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	customCtx := &model.Context{
		Data: g.Map{"Super": 0},
	}
	var err error
	customCtx.User, err = service.Gftoken().GetUser(ctx)

	if err != nil {
		g.Log().Error(ctx, err)
		response.JsonExit(r, 0, "获取用户信息失败")
	}
	service.Context().Init(r, customCtx)

	r.Middleware.Next()

}
func (s *sMiddleware) MCtx(r *ghttp.Request) {
	ctx := r.GetCtx()
	var (
		memberInfo *entity.Member
		err        error
	)
	memberInfo, err = service.Gftoken().GetMember(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		response.JsonExit(r, 0, "获取用户信息失败")
	}
	service.Context().MInit(r, memberInfo)

	r.Middleware.Next()
}
func (s *sMiddleware) MLog(r *ghttp.Request) {
	ctx := r.GetCtx()
	g.Dump(r.GetUrl(), service.Context().MGet(ctx).Id)
	r.Middleware.Next()
}
func (s *sMiddleware) Cas(r *ghttp.Request) {
	ctx := r.Context()
	userId := service.Context().Get(ctx).User.Id

	service.User().NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == userId {

			service.Context().SetData(ctx, g.Map{"Super": 1})
			return false
		}
		return true
	})

	if service.Context().Get(ctx).Data["Super"] == 1 {

		r.Middleware.Next()
		//不要再往后面执行
		return
	}
	apiList, err := service.Menu().GetNode(ctx, 2)
	var apiNode *entity.SysAuthRule
	url := gstr.TrimLeft(r.Request.URL.Path, "/")
	if err != nil {
		g.Log().Error(ctx, err)
		response.JsonExit(r, 0, "获取接口权限失败")
	}
	for _, m := range apiList {
		ms := gstr.SubStr(m.Name, 0, gstr.Pos(m.Name, "?"))
		if m.Name == url || ms == url {
			apiNode = m
			break
		}
	}
	//只验证存在数据库中的规则
	if apiNode != nil {
		if gstr.Equal(apiNode.Condition, "nocheck") {
			r.Middleware.Next()
			return
		}
		nodeId := apiNode.Id
		hasAccess := false
		hasAccess, err = service.Cas().CheckAuth(userId, nodeId, "")

		if err != nil {
			g.Log().Error(ctx, err)
			response.JsonExit(r, 50, "权限判断失败")
		}
		if !hasAccess {
			response.JsonExit(r, 50, "没有相应权限")
		}

	} else {
		response.JsonExit(r, 50, "没有设置权限")
	}
	r.Middleware.Next()
}
