// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberDao is the data access object for table member.
type MemberDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns MemberColumns // columns contains all the column names of Table for convenient usage.
}

// MemberColumns defines and stores column names for table member.
type MemberColumns struct {
	Id          string //
	Username    string //
	Realname    string //
	Nickname    string //
	Idcard      string //
	Group       string //
	Bigclass    string //
	Smallclass  string //
	Mobile      string //
	WebAuth     string //
	Pwd         string //
	Avatar      string //
	Sex         string // 0woman1man
	VerifyType  string //
	VerifyPhoto string //
	VerifyTime  string //
	Status      string // 2dongjie
	Openid      string //
}

// memberColumns holds the columns for table member.
var memberColumns = MemberColumns{
	Id:          "id",
	Username:    "username",
	Realname:    "realname",
	Nickname:    "nickname",
	Idcard:      "idcard",
	Group:       "group",
	Bigclass:    "bigclass",
	Smallclass:  "smallclass",
	Mobile:      "mobile",
	WebAuth:     "web_auth",
	Pwd:         "pwd",
	Avatar:      "avatar",
	Sex:         "sex",
	VerifyType:  "verify_type",
	VerifyPhoto: "verify_photo",
	VerifyTime:  "verify_time",
	Status:      "status",
	Openid:      "openid",
}

// NewMemberDao creates and returns a new DAO object for table data access.
func NewMemberDao() *MemberDao {
	return &MemberDao{
		group:   "default",
		table:   "member",
		columns: memberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MemberDao) Columns() MemberColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
