// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MfileDao is the data access object for table mfile.
type MfileDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns MfileColumns // columns contains all the column names of Table for convenient usage.
}

// MfileColumns defines and stores column names for table mfile.
type MfileColumns struct {
	Id        string // 自增ID
	Name      string // 文件名称
	Src       string // 本地文件存储路径
	Url       string // URL地址，可能为空
	MemberId  string // 操作用户
	CreatedAt string //
}

// mfileColumns holds the columns for table mfile.
var mfileColumns = MfileColumns{
	Id:        "id",
	Name:      "name",
	Src:       "src",
	Url:       "url",
	MemberId:  "member_id",
	CreatedAt: "created_at",
}

// NewMfileDao creates and returns a new DAO object for table data access.
func NewMfileDao() *MfileDao {
	return &MfileDao{
		group:   "default",
		table:   "mfile",
		columns: mfileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MfileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MfileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MfileDao) Columns() MfileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MfileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MfileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MfileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
