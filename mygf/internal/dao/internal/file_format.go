// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileFormatDao is the data access object for table file_format.
type FileFormatDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns FileFormatColumns // columns contains all the column names of Table for convenient usage.
}

// FileFormatColumns defines and stores column names for table file_format.
type FileFormatColumns struct {
	Id             string // 逻辑ID
	SystemFileUuid string // 文件uuid，system_file表的uuid
	Progress       string // 进度
	Path           string // 文件路径
	Status         string // 转码状态，0未开始，1转换中，2成功，3失败，4待确认
	Type           string // 类型，0=》无，1=》doc to pdf
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// fileFormatColumns holds the columns for table file_format.
var fileFormatColumns = FileFormatColumns{
	Id:             "id",
	SystemFileUuid: "system_file_uuid",
	Progress:       "progress",
	Path:           "path",
	Status:         "status",
	Type:           "type",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewFileFormatDao creates and returns a new DAO object for table data access.
func NewFileFormatDao() *FileFormatDao {
	return &FileFormatDao{
		group:   "default",
		table:   "file_format",
		columns: fileFormatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FileFormatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FileFormatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FileFormatDao) Columns() FileFormatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FileFormatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FileFormatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FileFormatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
