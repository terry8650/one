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
	IPersonal interface {
		ChangePwd(ctx context.Context, req *sys.PersonalChangePwdReq) (res *sys.PersonalChangePwdRes, err error)
		Logout(ctx context.Context, req *sys.PersonalLogoutReq) (res *sys.PersonalLogoutRes, err error)
	}
)

var (
	localPersonal IPersonal
)

func Personal() IPersonal {
	if localPersonal == nil {
		panic("implement not found for interface IPersonal, forgot register?")
	}
	return localPersonal
}

func RegisterPersonal(i IPersonal) {
	localPersonal = i
}