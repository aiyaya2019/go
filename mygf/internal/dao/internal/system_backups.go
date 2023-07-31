// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemBackupsDao is the data access object for table system_backups.
type SystemBackupsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SystemBackupsColumns // columns contains all the column names of Table for convenient usage.
}

// SystemBackupsColumns defines and stores column names for table system_backups.
type SystemBackupsColumns struct {
	Id         string // 主键
	Name       string // 备份名字
	CreateTime string // 创建时间
	Time       string // 备份花费时间
	FileSize   string // 文件大小
	Type       string // 1=>系统配置 2=>服务器文件
	Status     string // 1=>备份中 2=>成功 3=>失败
	Path       string // 文件路径
	SqlPath    string // 数据库文件存放路径
	Mark       string // 备注
	CurrentUse string // 0=>默认不使用1=>升级中 2=>成功使用 3=>升级失败
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// systemBackupsColumns holds the columns for table system_backups.
var systemBackupsColumns = SystemBackupsColumns{
	Id:         "id",
	Name:       "name",
	CreateTime: "create_time",
	Time:       "time",
	FileSize:   "file_size",
	Type:       "type",
	Status:     "status",
	Path:       "path",
	SqlPath:    "sql_path",
	Mark:       "mark",
	CurrentUse: "current_use",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSystemBackupsDao creates and returns a new DAO object for table data access.
func NewSystemBackupsDao() *SystemBackupsDao {
	return &SystemBackupsDao{
		group:   "default",
		table:   "system_backups",
		columns: systemBackupsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemBackupsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemBackupsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemBackupsDao) Columns() SystemBackupsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemBackupsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemBackupsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemBackupsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
