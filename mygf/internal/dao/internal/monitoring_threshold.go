// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MonitoringThresholdDao is the data access object for table monitoring_threshold.
type MonitoringThresholdDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns MonitoringThresholdColumns // columns contains all the column names of Table for convenient usage.
}

// MonitoringThresholdColumns defines and stores column names for table monitoring_threshold.
type MonitoringThresholdColumns struct {
	Id              string // 主键
	HarddiskPercent string // 磁盘使用率
	MemoryPercent   string // 内存使用率
	CpuPercent      string // CPU百分比
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// monitoringThresholdColumns holds the columns for table monitoring_threshold.
var monitoringThresholdColumns = MonitoringThresholdColumns{
	Id:              "id",
	HarddiskPercent: "harddisk_percent",
	MemoryPercent:   "memory_percent",
	CpuPercent:      "cpu_percent",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewMonitoringThresholdDao creates and returns a new DAO object for table data access.
func NewMonitoringThresholdDao() *MonitoringThresholdDao {
	return &MonitoringThresholdDao{
		group:   "default",
		table:   "monitoring_threshold",
		columns: monitoringThresholdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MonitoringThresholdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MonitoringThresholdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MonitoringThresholdDao) Columns() MonitoringThresholdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MonitoringThresholdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MonitoringThresholdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MonitoringThresholdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
