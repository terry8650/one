package user

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"one2.3/api/v1/sys"
	"one2.3/internal/consts"
	"one2.3/internal/dao"
	"one2.3/internal/model/do"
	"one2.3/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}
func New() *sUser {
	return &sUser{}
}

func (s *sUser) NotCheckAuthAdminIds(ctx context.Context) *gset.Set {
	ids := g.Cfg().MustGet(ctx, "system.notCheckAdminIds")
	var notCheckAuthAdminIds *gset.Set
	if !g.IsNil(ids) {
		notCheckAuthAdminIds = gset.NewFrom(ids)
	}
	return notCheckAuthAdminIds
}
func (s *sUser) List(ctx context.Context, req *sys.UserSearchReq) (res *sys.UserSearchRes, err error) {
	res = &sys.UserSearchRes{}
	m := dao.SysUser.Ctx(ctx)
	if req != nil {
		if req.UserName != "" {
			m = m.WhereLike(dao.SysUser.Columns().UserName, "%"+req.UserName+"%")
		}
		if req.Mobile != "" {
			m = m.WhereLike(dao.SysUser.Columns().Mobile, "%"+req.Mobile+"%")
		}
		if req.Status != "" {
			m = m.Where(dao.SysUser.Columns().UserStatus, req.Status)
		}
		if req.DeptId != "" {
			m = m.Where(dao.SysUser.Columns().DeptId, req.DeptId)
		}
		if req.PostId != "" {
			m = m.Where(dao.SysUser.Columns().PostId, req.PostId)
		}
	}
	res.Total, err = m.Count()
	if err != nil {
		return
	}

	res.Message = consts.OK
	res.Code = 0
	err = m.Page(req.PageNum, req.PageSize).Order("id desc").Scan(&res.UserList)

	return
}
func (s *sUser) Add(ctx context.Context, req *sys.UserAddReq) (err error) {
	salt := grand.S(10)
	pass := gmd5.MustEncryptString(gmd5.MustEncryptString("123456") + gmd5.MustEncryptString(salt))
	var userId int64
	userId, err = dao.SysUser.Ctx(ctx).InsertAndGetId(do.SysUser{
		UserName:     req.UserName,
		Mobile:       req.Mobile,
		UserNickname: req.UserNickname,
		UserPassword: pass,
		UserSalt:     salt,
		UserStatus:   req.UserStatus,
		UserEmail:    req.UserEmail,
		Sex:          req.Sex,
		DeptId:       req.DeptId,
		Remark:       req.Remark,
		PostId:       req.PostId,
	})
	if err != nil {
		return
	}
	_, err = service.Cas().AddUserRoles(gconv.Uint64(userId), req.RoleIds)

	return
}
func (s *sUser) Update(ctx context.Context, req *sys.UserUpdateReq) (err error) {
	_, err = dao.SysUser.Ctx(ctx).WherePri(req.Id).Update(do.SysUser{
		UserName:     req.UserName,
		Mobile:       req.Mobile,
		UserNickname: req.UserNickname,
		UserStatus:   req.UserStatus,
		UserEmail:    req.UserEmail,
		Sex:          req.Sex,
		DeptId:       req.DeptId,
		Remark:       req.Remark,
		PostId:       req.PostId,
	})
	if err != nil {
		return
	}
	_, err = service.Cas().DeleteUserRoles(req.Id)
	if err != nil {
		return
	}
	_, err = service.Cas().AddUserRoles(req.Id, req.RoleIds)
	return
}
func (s *sUser) Delete(ctx context.Context, req *sys.UserDeleteReq) (err error) {
	_, err = dao.SysUser.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return
	}
	_, err = service.Cas().DeleteUserRoles(req.Id)
	return
}
func (s *sUser) ChangePwd(ctx context.Context, req *sys.UserChangePwdReq) (err error) {
	pass := gmd5.MustEncryptString(gmd5.MustEncryptString(req.Password) + gmd5.MustEncryptString(req.Salt))
	_, err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, req.UserName).Update(do.SysUser{UserPassword: pass})
	if err != nil {
		return
	}
	service.Gftoken().DownUser(ctx, req.UserName)
	return
}
