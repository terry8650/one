package sys

import (
	"context"
	"one2.3/api/v1/sys"
	"one2.3/internal/service"
)

type cPersonal struct{}

var Personal = cPersonal{}

func (c *cPersonal) ChangePwd(ctx context.Context, req *sys.PersonalChangePwdReq) (res *sys.PersonalChangePwdRes, err error) {
	_, err = service.Personal().ChangePwd(ctx, req)
	return nil, err
}
func (c *cPersonal) Logout(ctx context.Context, req *sys.PersonalLogoutReq) (res *sys.PersonalLogoutRes, err error) {
	res, err = service.Personal().Logout(ctx, req)
	return
}
