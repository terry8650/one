// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAuthRuleDao is the data access object for table sys_auth_rule.
type SysAuthRuleDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SysAuthRuleColumns // columns contains all the column names of Table for convenient usage.
}

// SysAuthRuleColumns defines and stores column names for table sys_auth_rule.
type SysAuthRuleColumns struct {
	Id        string //
	Pid       string // 父ID
	Name      string // 规则名称
	Jump      string //
	Title     string // 规则名称
	Icon      string // 图标
	SmallAuth string // 接口细化权限说明，用来写入casbin第二个权限
	Condition string // 条件：nocheck就不用检测权限
	Remark    string // 备注
	Type      string // 类型1菜单 2接口
	Sort      string // 权重
	IsHide    string // 显示状态
	CreatedAt string // 创建日期
	UpdatedAt string // 修改日期
}

// sysAuthRuleColumns holds the columns for table sys_auth_rule.
var sysAuthRuleColumns = SysAuthRuleColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Jump:      "jump",
	Title:     "title",
	Icon:      "icon",
	SmallAuth: "small_auth",
	Condition: "condition",
	Remark:    "remark",
	Type:      "type",
	Sort:      "sort",
	IsHide:    "is_hide",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysAuthRuleDao creates and returns a new DAO object for table data access.
func NewSysAuthRuleDao() *SysAuthRuleDao {
	return &SysAuthRuleDao{
		group:   "default",
		table:   "sys_auth_rule",
		columns: sysAuthRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAuthRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAuthRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAuthRuleDao) Columns() SysAuthRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAuthRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAuthRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAuthRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}