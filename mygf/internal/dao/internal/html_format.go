// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HtmlFormatDao is the data access object for table html_format.
type HtmlFormatDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns HtmlFormatColumns // columns contains all the column names of Table for convenient usage.
}

// HtmlFormatColumns defines and stores column names for table html_format.
type HtmlFormatColumns struct {
	Id          string // 逻辑ID
	MeetingUuid string // 会议uuid，meeting表的uuid
	MultiUuid   string // 用户/标语uuid，user表的uuid/meeting_banner的uuid
	Url         string // url路径
	Path        string // 文件路径
	Progress    string // 进度
	Status      string // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Type        string // 类型，0=》无，1=》html to image，2=》标语
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// htmlFormatColumns holds the columns for table html_format.
var htmlFormatColumns = HtmlFormatColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	MultiUuid:   "multi_uuid",
	Url:         "url",
	Path:        "path",
	Progress:    "progress",
	Status:      "status",
	Type:        "type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewHtmlFormatDao creates and returns a new DAO object for table data access.
func NewHtmlFormatDao() *HtmlFormatDao {
	return &HtmlFormatDao{
		group:   "default",
		table:   "html_format",
		columns: htmlFormatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HtmlFormatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HtmlFormatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HtmlFormatDao) Columns() HtmlFormatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HtmlFormatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HtmlFormatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HtmlFormatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
