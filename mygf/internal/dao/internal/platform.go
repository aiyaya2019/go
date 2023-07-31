// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlatformDao is the data access object for table platform.
type PlatformDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns PlatformColumns // columns contains all the column names of Table for convenient usage.
}

// PlatformColumns defines and stores column names for table platform.
type PlatformColumns struct {
	Id        string //
	Uuid      string // 平台uuid
	Name      string // 平台名称
	IsLocal   string // 是否为本地：0远端；1本地
	Sysip     string // 平台ip
	Sysport   string // 平台端口
	Sysid     string // 平台自定义id，唯一
	Status    string // 状态:0离线，1=>在线
	Enable    string // 0禁用；1开启
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// platformColumns holds the columns for table platform.
var platformColumns = PlatformColumns{
	Id:        "id",
	Uuid:      "uuid",
	Name:      "name",
	IsLocal:   "is_local",
	Sysip:     "sysip",
	Sysport:   "sysport",
	Sysid:     "sysid",
	Status:    "status",
	Enable:    "enable",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPlatformDao creates and returns a new DAO object for table data access.
func NewPlatformDao() *PlatformDao {
	return &PlatformDao{
		group:   "default",
		table:   "platform",
		columns: platformColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PlatformDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PlatformDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PlatformDao) Columns() PlatformColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PlatformDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PlatformDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PlatformDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
