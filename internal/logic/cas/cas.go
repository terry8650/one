package cas

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
	"one2.3/internal/service"
	casbin "one2.3/utility/casbin"
)

type sCas struct {
	Enforce *casbin.Enforcer
}

func init() {

	service.RegisterCas(New())
}
func New() *sCas {
	db := g.DB()
	var ctx = gctx.New()
	config, _ := g.Cfg().Get(ctx, "casbin")
	var options casbin.Options

	confErr := gconv.Scan(config, &options)
	if confErr != nil {
		fmt.Println("init casbin config error")
	}

	options.DB = db
	e, err := casbin.NewEnforcer(&options)
	if err != nil {
		log.Fatalf("Casbin init failure:%s \n", err.Error())
	}
	return &sCas{e}
}

func (s *sCas) CheckAuth(Uid uint64, Node uint, Auth string) (bool, error) {
	userStr := "u_" + gconv.String(Uid)

	return s.Enforce.Enforce(userStr, gconv.String(Node), Auth)
}

func (s *sCas) AddPolicy(role uint, nodeId uint) (bool, error) {
	roleStr := "r_" + gconv.String(role)
	return s.Enforce.AddPolicy(roleStr, nodeId, "d")
}

// AddPolicies 批量添加角色对应节点，第四个参数d为默认
func (s *sCas) AddPolicies(ctx context.Context, role uint, nodeIds []uint) (bool, error) {
	listNum := len(nodeIds)

	ruleList := make([][]string, listNum)
	for key, val := range nodeIds {
		ruleList[key] = []string{"r_" + gconv.String(role), gconv.String(val), "d"}

	}

	return s.Enforce.AddPolicies(ruleList)
}
func (s *sCas) AddPoliciesMore(newRule map[int][]string) (added bool, err error) {
	listLength := len(newRule)
	news := make([][]string, listLength)
	for i, v := range newRule {
		news[i] = v
	}
	added, err = s.Enforce.AddPolicies(news)
	return
}
func (s *sCas) AddUserRoles(user uint64, roleIds []uint) (bool, error) {
	listLength := len(roleIds)
	roles := make([][]string, listLength)
	for key, val := range roleIds {
		roles[key] = []string{"u_" + gconv.String(user), "r_" + gconv.String(val)}

	}
	return s.Enforce.AddGroupingPolicies(roles)

}
func (s *sCas) AddRoleUsers(role uint64, userIds []uint) (bool, error) {
	listLength := len(userIds)
	users := make([][]string, listLength)
	for key, val := range userIds {
		users[key] = []string{"u_" + gconv.String(val), "r_" + gconv.String(role)}

	}
	return s.Enforce.AddGroupingPolicies(users)

}

// DeleteRoleNodes 删除角色对应的所有节点
func (s *sCas) DeleteRoleNodes(role uint) (res bool, err error) {
	res, err = s.Enforce.DeletePermissionsForUser("r_" + gconv.String(role))

	return
}
func (s *sCas) GetRoleNodes(role uint) (res [][]string) {
	res = s.Enforce.GetPermissionsForUser("r_" + gconv.String(role))

	return
}
func (s *sCas) GetUserRoles(UserId uint64) (roleIds []uint, err error) {
	var roles []string
	roles, err = s.Enforce.GetRolesForUser("u_" + gconv.String(UserId))
	listLength := len(roles)
	roleIds = make([]uint, listLength)
	for i, role := range roles {
		roleIds[i] = gconv.Uint(gstr.StrEx(role, "_"))
	}

	return
}

// DeleteNodeRoles 删除节点对应的所有角色，删除节点时用到
func (s *sCas) DeleteNodeRoles(node uint) (res bool, err error) {
	res, err = s.Enforce.DeletePermission(gconv.String(node))
	return
}

// DeleteUserRoles 删除用户对应的所有角色
func (s *sCas) DeleteUserRoles(user uint64) (res bool, err error) {
	res, err = s.Enforce.DeleteRolesForUser("u_" + gconv.String(user))

	return
}
func (s *sCas) DelOnePolicy(Rule []string) (updated bool, err error) {
	updated, err = s.Enforce.RemovePolicy(Rule)
	return
}
func (s *sCas) UpdatePolicy(oldRule, newRule []string) (updated bool, err error) {
	updated, err = s.Enforce.UpdatePolicy(oldRule, newRule)
	return
}
