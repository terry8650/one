// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"one2.3/api/v1/sys"
	"one2.3/internal/model"
	"one2.3/internal/model/entity"
)

type (
	IMenu interface {
		Add(ctx context.Context, req *sys.MenuAddReq) (err error)
		Edit(ctx context.Context, req *sys.MenuEditReq) (err error)
		Delete(ctx context.Context, id uint) (err error)
		GetMenuList(ctx context.Context, req *sys.MenuSearchReq) (res *sys.MenuSearchRes, err error)
		GetAuthList(ctx context.Context) ([]*sys.MenuAuthList, error)
		GetNode(ctx context.Context, nodeType uint) ([]*entity.SysAuthRule, error)
		GetAllNode(ctx context.Context) (node []*entity.SysAuthRule, err error)
		GetMenuListTree(pid uint, list []*model.MenuInfoRes) []*model.SysMenuTreeRes
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}