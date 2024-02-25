package personal

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"one2.3/api/v1/sys"
	"one2.3/internal/dao"
	"one2.3/internal/model"
	"one2.3/internal/model/do"
	"one2.3/internal/service"
)

type sPersonal struct {
}

func init() {
	service.RegisterPersonal(New())
}
func New() *sPersonal {
	return &sPersonal{}
}
func (s *sPersonal) ChangePwd(ctx context.Context, req *sys.PersonalChangePwdReq) (res *sys.PersonalChangePwdRes, err error) {

	_, err = service.Gftoken().CheckUserPassword(ctx, model.Check{
		Username: service.Context().Get(ctx).User.UserName,
		Password: req.OldPassWord,
	})
	if err != nil {
		return nil, err
	}
	pass := gmd5.MustEncryptString(gmd5.MustEncryptString(req.Password) + gmd5.MustEncryptString(service.Context().Get(ctx).User.UserSalt))
	_, err = dao.SysUser.Ctx(ctx).WherePri(service.Context().Get(ctx).User.Id).Update(do.SysUser{UserPassword: pass})
	if err != nil {
		return nil, err
	}
	service.Gftoken().DownUser(ctx, service.Context().Get(ctx).User.UserName)
	return
}
func (s *sPersonal) Logout(ctx context.Context, req *sys.PersonalLogoutReq) (res *sys.PersonalLogoutRes, err error) {
	ye := service.Gftoken().DownUser(ctx, service.Context().Get(ctx).User.UserName)
	if !ye {
		return nil, gerror.New("登出失败")
	}
	return nil, nil
}
