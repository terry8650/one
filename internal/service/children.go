// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IChildren interface {
		CheckClassMems(ctx context.Context, cid string, mid string)
	}
)

var (
	localChildren IChildren
)

func Children() IChildren {
	if localChildren == nil {
		panic("implement not found for interface IChildren, forgot register?")
	}
	return localChildren
}

func RegisterChildren(i IChildren) {
	localChildren = i
}