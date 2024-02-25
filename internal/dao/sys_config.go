// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"one2.3/internal/dao/internal"
)

// internalSysConfigDao is internal type for wrapping internal DAO implements.
type internalSysConfigDao = *internal.SysConfigDao

// sysConfigDao is the data access object for table sys_config.
// You can define custom methods on it to extend its functionality as you wish.
type sysConfigDao struct {
	internalSysConfigDao
}

var (
	// SysConfig is globally public accessible object for table sys_config operations.
	SysConfig = sysConfigDao{
		internal.NewSysConfigDao(),
	}
)

// Fill with you ideas below.
