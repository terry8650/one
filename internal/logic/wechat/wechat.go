package wechat

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/work"
	"github.com/silenceper/wechat/v2/work/config"
	"net/url"
	"one2.3/internal/model"
	"one2.3/internal/service"
)

type sWeChat struct {
	*work.Work
}

func New() *sWeChat {
	var (
		weConf  *config.Config
		weRedis *cache.RedisOpts
	)
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, "wework").Scan(&weConf)
	if err != nil {
		return nil
	}

	err = g.Cfg().MustGet(ctx, "weredis").Scan(&weRedis)
	if err != nil {
		return nil
	}

	weConf.Cache = cache.NewRedis(ctx, weRedis)

	wework := wechat.NewWechat().GetWork(weConf)

	return &sWeChat{Work: wework}
}
func init() {
	service.RegisterWeChat(New())
}
func (s *sWeChat) LoginUrl(ctx context.Context, tag string) (uurl string) {

	uurl = s.GetOauth().GetTargetURL("http://go.nanpi.site/pub/wechat/callback?tag=" + url.QueryEscape(tag))

	return uurl
}
func (s *sWeChat) GetOpenId(ctx context.Context, code string) (openid string, err error) {
	fromCode, err := s.GetOauth().UserFromCode(code)
	if err != nil {
		return
	}

	if fromCode.OpenID == "" {
		fromCode.OpenID, err = s.GetAddressList().ConvertToOpenID(fromCode.UserID)
	}

	return fromCode.OpenID, err
}

func (s *sWeChat) SendMsg(ctx context.Context, msg *model.MsgInfo) (err error) {

	v := model.AllMsg{
		ToUser: msg.ToUser,

		Msgtype: "text",
		Text:    &model.MsgType{Content: msg.Content},
		AgentId: msg.AgentId,
	}
	_, err = s.Work.GetMessage().Send("os", v)
	return
}
