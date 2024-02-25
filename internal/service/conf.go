// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"one2.3/api/v1/sys"
)

type (
	IConf interface {
		List(ctx context.Context) (res *sys.ConfListRes, err error)
		GetVal(ctx context.Context, req *sys.ConfGetReq) (res *sys.ConfGetRes, err error)
		Add(ctx context.Context, req *sys.ConfAddReq) (err error)
		Update(ctx context.Context, req *sys.ConfUpdateReq) (err error)
		Delete(ctx context.Context, ids []uint) (err error)
	}
)

var (
	localConf IConf
)

func Conf() IConf {
	if localConf == nil {
		panic("implement not found for interface IConf, forgot register?")
	}
	return localConf
}

func RegisterConf(i IConf) {
	localConf = i
}
