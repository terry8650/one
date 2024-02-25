// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
)

type (
	ICache interface {
		Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error)
		Get(ctx context.Context, key string) (*gvar.Var, error)
		Contains(ctx context.Context, key string) bool
		Remove(ctx context.Context, key string) (lastValue interface{})
		GetOrSetFuncLock(ctx context.Context, key string, f gcache.Func, duration time.Duration) (*gvar.Var, error)
	}
)

var (
	localCache ICache
)

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
