// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MonitoringDao is the data access object for table monitoring.
type MonitoringDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns MonitoringColumns // columns contains all the column names of Table for convenient usage.
}

// MonitoringColumns defines and stores column names for table monitoring.
type MonitoringColumns struct {
	Id               string //
	HarddiskUse      string // 硬盘使用空间
	Harddisk         string // 硬盘
	MemoryUse        string // 内存使用
	Memory           string // 内存
	CpuPercent       string // cpu使用率
	NetworkUpspeed   string // 上行网速
	NetworkDownspeed string // 下行网速
	TerminalMac      string // 终端mac，默认1为8300服务器
	LastViewTime     string // 最后一次浏览通知时间
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// monitoringColumns holds the columns for table monitoring.
var monitoringColumns = MonitoringColumns{
	Id:               "id",
	HarddiskUse:      "harddisk_use",
	Harddisk:         "harddisk",
	MemoryUse:        "memory_use",
	Memory:           "memory",
	CpuPercent:       "cpu_percent",
	NetworkUpspeed:   "network_upspeed",
	NetworkDownspeed: "network_downspeed",
	TerminalMac:      "terminal_mac",
	LastViewTime:     "last_view_time",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewMonitoringDao creates and returns a new DAO object for table data access.
func NewMonitoringDao() *MonitoringDao {
	return &MonitoringDao{
		group:   "default",
		table:   "monitoring",
		columns: monitoringColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MonitoringDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MonitoringDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MonitoringDao) Columns() MonitoringColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MonitoringDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MonitoringDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MonitoringDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
