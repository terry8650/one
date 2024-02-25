package menu

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"one2.3/api/v1/sys"
	"one2.3/internal/consts"
	"one2.3/internal/dao"
	"one2.3/internal/model"
	"one2.3/internal/model/do"
	"one2.3/internal/model/entity"
	"one2.3/internal/service"
)

type sMenu struct {
}

func init() {
	service.RegisterMenu(New())
}
func New() *sMenu {
	return &sMenu{}
}
func (s *sMenu) Add(ctx context.Context, req *sys.MenuAddReq) (err error) {

	_, err = dao.SysAuthRule.Ctx(ctx).Insert(do.SysAuthRule{
		Title:     req.Title,
		Name:      req.Name,
		Jump:      req.Jump,
		Icon:      req.Icon,
		Pid:       req.Pid,
		Remark:    req.Remark,
		Type:      req.Type,
		Condition: req.Condition,
		SmallAuth: req.SmallAuth,
	})
	if err != nil {
		return
	}
	service.Cache().Remove(ctx, consts.SysNode)

	return
}
func (s *sMenu) Edit(ctx context.Context, req *sys.MenuEditReq) (err error) {

	_, err = dao.SysAuthRule.Ctx(ctx).WherePri(req.Id).Update(do.SysAuthRule{
		Title:     req.Title,
		Name:      req.Name,
		Jump:      req.Jump,
		Icon:      req.Icon,
		Pid:       req.Pid,
		Remark:    req.Remark,
		Type:      req.Type,
		Condition: req.Condition,
		SmallAuth: req.SmallAuth,
	})
	if err != nil {
		return
	}
	service.Cache().Remove(ctx, consts.SysNode)

	return
}
func (s *sMenu) Delete(ctx context.Context, id uint) (err error) {
	if s.existSon(ctx, id) {
		err = gerror.New("先删除子节点")
		return
	}
	_, err = dao.SysAuthRule.Ctx(ctx).Where(dao.SysAuthRule.Columns().Id, id).Delete()
	if err != nil {
		return
	}

	service.Cache().Remove(ctx, consts.SysNode)
	_, err = service.Cas().DeleteNodeRoles(id)

	return
}
func (s *sMenu) existSon(ctx context.Context, id uint) bool {

	d, err := dao.SysAuthRule.Ctx(ctx).Where(dao.SysAuthRule.Columns().Pid, id).One()

	if err != nil {
		g.Log().Error(ctx, err)
		return false
	}

	return !d.IsEmpty()
}
func (s *sMenu) GetMenuList(ctx context.Context, req *sys.MenuSearchReq) (res *sys.MenuSearchRes, err error) {
	res = &sys.MenuSearchRes{}
	m := dao.SysAuthRule.Ctx(ctx)
	if req.Title != "" {
		m = m.WhereLike(dao.SysAuthRule.Columns().Title, "%"+req.Title+"%")
	}
	if req.Name != "" {
		m = m.WhereLike(dao.SysAuthRule.Columns().Name, "%"+req.Name+"%")
	}
	res.Total, err = m.Count()
	if err != nil {
		return
	}

	res.Message = consts.OK
	res.Code = 0
	err = m.Fields(entity.SysAuthRule{}).Order("id asc").Scan(&res.MenuList)

	return
}
func (s *sMenu) GetAuthList(ctx context.Context) ([]*sys.MenuAuthList, error) {

	all, err := s.GetNode(ctx, 1)
	res := make([]*sys.MenuAuthList, 0, len(all))
	if service.Context().Get(ctx).Data["Super"] == 1 {
		for _, rule := range all {

			res = append(res, &sys.MenuAuthList{SysAuthRule: rule, MyAuth: "admin"})

		}
		return res, err
	} else {
		var (
			rIds []uint
		)
		menuIds := map[uint][]string{}
		rIds, err = service.Cas().GetUserRoles(service.Context().Get(ctx).User.Id)

		for _, id := range rIds {
			rns := service.Cas().GetRoleNodes(id)

			for _, rn := range rns {
				mid := gconv.Uint(rn[1])
				if _, ok := menuIds[mid]; ok {
					menuIds[mid] = []string{rn[1], menuIds[mid][1] + "," + rn[2]}
				} else {
					menuIds[mid] = []string{rn[1], rn[2]}
				}

			}

		}
		for _, v := range all {
			if _, ok := menuIds[gconv.Uint(v.Id)]; gstr.Equal(v.Condition, "nocheck") || ok {
				res = append(res, &sys.MenuAuthList{SysAuthRule: v, MyAuth: menuIds[gconv.Uint(v.Id)][1]})
			}

		}

	}

	return res, err
}
func (s *sMenu) GetNode(ctx context.Context, nodeType uint) ([]*entity.SysAuthRule, error) {
	list, err := s.GetAllNode(ctx)
	if err != nil {
		return nil, err
	}
	gList := make([]*entity.SysAuthRule, 0, len(list))
	for _, v := range list {
		if v.Type == nodeType {
			gList = append(gList, v)
		}
	}
	return gList, nil
}
func (s *sMenu) GetAllNode(ctx context.Context) (node []*entity.SysAuthRule, err error) {
	var iNode interface{}
	iNode, err = service.Cache().GetOrSetFuncLock(ctx, consts.SysNode, func(ctx context.Context) (value interface{}, err error) {
		err = dao.SysAuthRule.Ctx(ctx).
			Fields(entity.SysAuthRule{}).Order("sort desc,id asc").Scan(&node)
		if err != nil {
			return nil, err
		}
		value = node
		return
	}, 0)
	if err != nil {
		return
	}
	if iNode != nil {
		err = gconv.Struct(iNode, &node)
	}
	return
}

func (s *sMenu) GetMenuListTree(pid uint, list []*model.MenuInfoRes) []*model.SysMenuTreeRes {
	tree := make([]*model.SysMenuTreeRes, 0, len(list))
	for _, menu := range list {
		if menu.Pid == pid {
			t := &model.SysMenuTreeRes{
				MenuInfoRes: menu,
			}
			child := s.GetMenuListTree(menu.Id, list)
			if child != nil {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return tree
}
