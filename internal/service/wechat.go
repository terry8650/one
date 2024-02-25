// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"one2.3/internal/model"
)

type (
	IWeChat interface {
		LoginUrl(ctx context.Context, tag string) (uurl string)
		GetOpenId(ctx context.Context, code string) (openid string, err error)
		SendMsg(ctx context.Context, msg *model.MsgInfo) (err error)
	}
)

var (
	localWeChat IWeChat
)

func WeChat() IWeChat {
	if localWeChat == nil {
		panic("implement not found for interface IWeChat, forgot register?")
	}
	return localWeChat
}

func RegisterWeChat(i IWeChat) {
	localWeChat = i
}
