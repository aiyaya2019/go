// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideoFormatDao is the data access object for table video_format.
type VideoFormatDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns VideoFormatColumns // columns contains all the column names of Table for convenient usage.
}

// VideoFormatColumns defines and stores column names for table video_format.
type VideoFormatColumns struct {
	Id             string // 逻辑ID
	SystemFileUuid string // 视频文件uuid，system_file表的uuid
	Progress       string // 进度
	Status         string // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败，4 =》待确认
	Path           string // 文件路径
	Type           string // 0=》 转MP4，aac，1=》全部 copy，2=》视频copy，音频转aac
	Force          string // 强制转码0:否1:是
	Remain         string // 预计剩余时长
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// videoFormatColumns holds the columns for table video_format.
var videoFormatColumns = VideoFormatColumns{
	Id:             "id",
	SystemFileUuid: "system_file_uuid",
	Progress:       "progress",
	Status:         "status",
	Path:           "path",
	Type:           "type",
	Force:          "force",
	Remain:         "remain",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewVideoFormatDao creates and returns a new DAO object for table data access.
func NewVideoFormatDao() *VideoFormatDao {
	return &VideoFormatDao{
		group:   "default",
		table:   "video_format",
		columns: videoFormatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VideoFormatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VideoFormatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VideoFormatDao) Columns() VideoFormatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VideoFormatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VideoFormatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VideoFormatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
