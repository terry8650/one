package cache

import (
	"context"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"one2.3/internal/service"
	"time"
)

type sCache struct {
	CachePrefix string //缓存前缀
	cache       *gcache.Cache
}

func init() {
	service.RegisterCache(New())
}
func New() *sCache {
	ctx := gctx.New()
	model := g.Cfg().MustGet(ctx, "system.cache.model").String()
	prefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	ca := gcache.New()
	if model == "redis" {
		ca.SetAdapter(gcache.NewAdapterRedis(g.Redis()))

	}

	return &sCache{
		CachePrefix: prefix,
		cache:       ca,
	}
}
func (s *sCache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error) {
	err = s.cache.Set(ctx, s.CachePrefix+key, value, duration)
	return
}
func (s *sCache) Get(ctx context.Context, key string) (*gvar.Var, error) {
	return s.cache.Get(ctx, s.CachePrefix+key)
}
func (s *sCache) Contains(ctx context.Context, key string) bool {
	v, _ := s.cache.Contains(ctx, s.CachePrefix+key)
	return v
}
func (s *sCache) Remove(ctx context.Context, key string) (lastValue interface{}) {
	lastValue, _ = s.cache.Remove(ctx, s.CachePrefix+key)
	return
}

func (s *sCache) GetOrSetFuncLock(ctx context.Context, key string, f gcache.Func, duration time.Duration) (*gvar.Var, error) {
	return s.cache.GetOrSetFuncLock(ctx, s.CachePrefix+key, f, duration)
}
