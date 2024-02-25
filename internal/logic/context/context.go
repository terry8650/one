package context

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"one2.3/internal/consts"
	"one2.3/internal/model"
	"one2.3/internal/model/entity"
	"one2.3/internal/service"
)

type sContext struct{}

func init() {
	service.RegisterContext(New())
}

func New() *sContext {
	return &sContext{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.CtxKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.CtxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.CtxUser) {
	s.Get(ctx).User = ctxUser
}

// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetData(ctx context.Context, data g.Map) {
	s.Get(ctx).Data = data
}

// 以下是member的操作
func (s *sContext) MInit(r *ghttp.Request, customCtx *entity.Member) {
	r.SetCtxVar(consts.CtxMKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) MGet(ctx context.Context) *entity.Member {
	value := ctx.Value(consts.CtxMKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*entity.Member); ok {
		return localCtx
	}
	return nil
}
