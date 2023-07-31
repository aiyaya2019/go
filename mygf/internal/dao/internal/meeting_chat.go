// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingChatDao is the data access object for table meeting_chat.
type MeetingChatDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns MeetingChatColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingChatColumns defines and stores column names for table meeting_chat.
type MeetingChatColumns struct {
	Id          string // 逻辑ID
	MeetingUuid string // 会议uuid，meeting表的uuid
	UserUuid    string // 用户uuid，user表的uuid
	MessageType string // 消息类型：1=>聊天消息，2=> 系统消息
	Type        string // 消息内容类型0=>文字，1=> 图片，2=> 语音
	Content     string // 消息内容
	SendTime    string // 发送时间
	AcceptUuid  string // user表的uuid或platform表的uuid；0 代表全部uuid  多个uuid用英文半角,隔开
	Mark        string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// meetingChatColumns holds the columns for table meeting_chat.
var meetingChatColumns = MeetingChatColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	UserUuid:    "user_uuid",
	MessageType: "message_type",
	Type:        "type",
	Content:     "content",
	SendTime:    "send_time",
	AcceptUuid:  "accept_uuid",
	Mark:        "mark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMeetingChatDao creates and returns a new DAO object for table data access.
func NewMeetingChatDao() *MeetingChatDao {
	return &MeetingChatDao{
		group:   "default",
		table:   "meeting_chat",
		columns: meetingChatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingChatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingChatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingChatDao) Columns() MeetingChatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingChatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingChatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingChatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
