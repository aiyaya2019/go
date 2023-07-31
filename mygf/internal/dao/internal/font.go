// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FontDao is the data access object for table font.
type FontDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns FontColumns // columns contains all the column names of Table for convenient usage.
}

// FontColumns defines and stores column names for table font.
type FontColumns struct {
	Id         string // 逻辑ID
	Name       string // 字体名称
	Path       string // 文件路径
	UploadTime string // 上传时间
	Status     string // 状态;1:启用,0:禁用
	IsSysFont  string // 是否系统自带字体;1:是,0:否
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// fontColumns holds the columns for table font.
var fontColumns = FontColumns{
	Id:         "id",
	Name:       "name",
	Path:       "path",
	UploadTime: "upload_time",
	Status:     "status",
	IsSysFont:  "is_sys_font",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewFontDao creates and returns a new DAO object for table data access.
func NewFontDao() *FontDao {
	return &FontDao{
		group:   "default",
		table:   "font",
		columns: fontColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FontDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FontDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FontDao) Columns() FontColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FontDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FontDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FontDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
