// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingExtraDao is the data access object for table meeting_extra.
type MeetingExtraDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns MeetingExtraColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingExtraColumns defines and stores column names for table meeting_extra.
type MeetingExtraColumns struct {
	MeetingUuid     string // 会议uuid，meeting表的uuid
	MenuTab         string // 选择的菜单数值
	NameplateCustom string // 铭牌的自定义文本，限制50个汉字
	WelcomeCustom   string // 欢迎界面的自定义文本，限制50个汉字
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// meetingExtraColumns holds the columns for table meeting_extra.
var meetingExtraColumns = MeetingExtraColumns{
	MeetingUuid:     "meeting_uuid",
	MenuTab:         "menu_tab",
	NameplateCustom: "nameplate_custom",
	WelcomeCustom:   "welcome_custom",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewMeetingExtraDao creates and returns a new DAO object for table data access.
func NewMeetingExtraDao() *MeetingExtraDao {
	return &MeetingExtraDao{
		group:   "default",
		table:   "meeting_extra",
		columns: meetingExtraColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingExtraDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingExtraDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingExtraDao) Columns() MeetingExtraColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingExtraDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingExtraDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingExtraDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}