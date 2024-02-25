package cmd

import (
	"context"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/util/gmode"
	"one2.3/internal/consts"
	"one2.3/internal/controller/common"
	"one2.3/internal/controller/home"
	"one2.3/internal/controller/sys"
	"one2.3/internal/service"
	"one2.3/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			if uploadPath == "" {
				g.Log().Fatal(ctx, "文件上传配置路径不能为空")
			}

			s.AddStaticPath("/upload", uploadPath)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().ResponseHandler,
				)
				group.Group("/pub", func(group *ghttp.RouterGroup) {
					group.Bind(
						common.Captcha,
						sys.Publi,
						home.Wechat,
					)

				})

				group.Group("/sys", func(group *ghttp.RouterGroup) {
					err := service.Gftoken().Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Middleware(service.Middleware().Ctx)
					group.Bind(common.File, sys.Personal)
					group.Middleware(service.Middleware().Cas)
					group.Bind(
						sys.Post,
						sys.Monitor,
						sys.Menu,
						sys.Role,
						sys.Dept,
						sys.User,
						sys.Conf,
					)

				})
				group.Group("/home", func(group *ghttp.RouterGroup) {
					err := service.Gftoken().Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Middleware(service.Middleware().MCtx, service.Middleware().MLog)

					group.Bind(

						home.Menu,

						home.Member,
						home.MFile,
						home.Children,

						//common.File,
					)
				})
			})
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = response.JsonRes{}
	openapi.Config.CommonResponseDataField = `data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIContactName,
			URL:  consts.OpenAPIContactUrl,
		},
	}
}
