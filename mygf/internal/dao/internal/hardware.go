// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HardwareDao is the data access object for table hardware.
type HardwareDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns HardwareColumns // columns contains all the column names of Table for convenient usage.
}

// HardwareColumns defines and stores column names for table hardware.
type HardwareColumns struct {
	Id            string //
	CpuPercent    string // cpu使用率
	MemoryPercent string // 内存使用率
	HarddiskSpace string // 硬盘剩余空间
	TerminalMac   string // 终端mac，默认1为8300服务器
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// hardwareColumns holds the columns for table hardware.
var hardwareColumns = HardwareColumns{
	Id:            "id",
	CpuPercent:    "cpu_percent",
	MemoryPercent: "memory_percent",
	HarddiskSpace: "harddisk_space",
	TerminalMac:   "terminal_mac",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewHardwareDao creates and returns a new DAO object for table data access.
func NewHardwareDao() *HardwareDao {
	return &HardwareDao{
		group:   "default",
		table:   "hardware",
		columns: hardwareColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HardwareDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HardwareDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HardwareDao) Columns() HardwareColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HardwareDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HardwareDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HardwareDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
