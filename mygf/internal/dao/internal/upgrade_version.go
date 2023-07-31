// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UpgradeVersionDao is the data access object for table upgrade_version.
type UpgradeVersionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UpgradeVersionColumns // columns contains all the column names of Table for convenient usage.
}

// UpgradeVersionColumns defines and stores column names for table upgrade_version.
type UpgradeVersionColumns struct {
	Id         string //
	UserUuid   string // 上传用户uuid，user表的uuid
	Name       string // 文件名
	Versions   string // 版本
	Url        string // 文件地址
	Type       string // 4=>cms,5=>mysql,7=>ps,10=>大屏,11=>client,12=>守护
	CurrentUse string // cms、client、mysql使用版本
	Mark       string // 备注
	Deleted    string // 1=>删除
	Status     string // 1=>升级中 2=>成功 3=>失败
	OsType     string // 1：windows；2：android; 3: ios;4: kylin； 5：centos;
	CreatedAt  string // 创建时间/上传时间
	UpdatedAt  string // 更新时间
}

// upgradeVersionColumns holds the columns for table upgrade_version.
var upgradeVersionColumns = UpgradeVersionColumns{
	Id:         "id",
	UserUuid:   "user_uuid",
	Name:       "name",
	Versions:   "versions",
	Url:        "url",
	Type:       "type",
	CurrentUse: "current_use",
	Mark:       "mark",
	Deleted:    "deleted",
	Status:     "status",
	OsType:     "os_type",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewUpgradeVersionDao creates and returns a new DAO object for table data access.
func NewUpgradeVersionDao() *UpgradeVersionDao {
	return &UpgradeVersionDao{
		group:   "default",
		table:   "upgrade_version",
		columns: upgradeVersionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UpgradeVersionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UpgradeVersionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UpgradeVersionDao) Columns() UpgradeVersionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UpgradeVersionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UpgradeVersionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UpgradeVersionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
