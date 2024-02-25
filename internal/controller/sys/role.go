package sys

import (
	"context"
	"one2.3/api/v1/sys"
	"one2.3/internal/service"
	"one2.3/utility/response"
)

type cRole struct{}

var Role = cRole{}

func (c *cRole) List(ctx context.Context, req *sys.RoleListReq) (res *sys.RoleListRes, err error) {
	res, err = service.Role().List(ctx, req)
	if err == nil {
		response.JsonOri(ctx, res)
	}
	return
}
func (c *cRole) Add(ctx context.Context, req *sys.RoleAddReq) (res *sys.RoleAddRes, err error) {
	err = service.Role().Add(ctx, req)
	return
}
func (c *cRole) Update(ctx context.Context, req *sys.RoleUpdateReq) (res *sys.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, req)
	return
}
func (c *cRole) Delete(ctx context.Context, req *sys.RoleDeleteReq) (res *sys.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Ids)
	return
}
func (c *cRole) DelOneRule(ctx context.Context, req *sys.DelOneRuleReq) (res *sys.DelOneRuleRes, err error) {
	_, err = service.Cas().DelOnePolicy(req.Auth)
	return
}
func (c *cRole) GetRoleNodes(ctx context.Context, req *sys.RoleNodesReq) (res *sys.RoleNodesRes, err error) {
	nodes := service.Cas().GetRoleNodes(req.RoleId)

	res = new(sys.RoleNodesRes)
	res.Auth = nodes
	return res, nil
}
func (c *cRole) UpdateRoleNodes(ctx context.Context, req *sys.RoleNodesUpdateReq) (res *sys.RoleNodesUpdateRes, err error) {
	_, err = service.Cas().DeleteRoleNodes(req.RoleId)
	if err != nil {
		return
	}
	_, err = service.Cas().AddPolicies(ctx, req.RoleId, req.Nodes)

	return
}
func (c *cRole) UpdateOldRules(ctx context.Context, req *sys.UpdateOldRulesReq) (res *sys.UpdateOldRulesRes, err error) {
	_, err = service.Cas().DeleteRoleNodes(req.RoleId)
	if err != nil {
		return
	}
	_, err = service.Cas().AddPoliciesMore(req.NewAuth)
	return
}
func (c *cRole) UpdateSingleRule(ctx context.Context, req *sys.UpdateSingleRuleReq) (res *sys.UpdateSingleRuleRes, err error) {
	_, err = service.Cas().UpdatePolicy(req.OldAuth, req.NewAuth)
	return
}
