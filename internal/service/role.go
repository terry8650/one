// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"one2.3/api/v1/sys"
	"one2.3/internal/model/entity"
)

type (
	IRole interface {
		List(ctx context.Context, req *sys.RoleListReq) (res *sys.RoleListRes, err error)
		Add(ctx context.Context, req *sys.RoleAddReq) (err error)
		Update(ctx context.Context, req *sys.RoleUpdateReq) (err error)
		Delete(ctx context.Context, ids []int) (err error)
		GetRole(ctx context.Context) (list []*entity.SysRole, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
