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
	IDept interface {
		List(ctx context.Context) (res *sys.DeptListRes, err error)
		Add(ctx context.Context, req *sys.DeptAddReq) (err error)
		Update(ctx context.Context, req *sys.DeptUpdateReq) (err error)
		Delete(ctx context.Context, id int64) (err error)
	}
)

var (
	localDept IDept
)

func Dept() IDept {
	if localDept == nil {
		panic("implement not found for interface IDept, forgot register?")
	}
	return localDept
}

func RegisterDept(i IDept) {
	localDept = i
}
