// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingSignDao is the data access object for table meeting_sign.
type MeetingSignDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns MeetingSignColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingSignColumns defines and stores column names for table meeting_sign.
type MeetingSignColumns struct {
	Id             string // 逻辑ID
	MeetingUuid    string // 会议uuid，meeting表的uuid
	SystemFileUuid string // 文件uuid，system_file表的uuid
	SignStartTime  string // 发起签到时间
	Title          string // 标题
	SignType       string // 签到方式：1签名签到；2登录即签到 3按钮签到 4拍照签到 5人脸签到 6指纹签到
	Status         string // 当前签到状态：0未开始，1签到中, 2 已结束
	QrCode         string // 签到二维码
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// meetingSignColumns holds the columns for table meeting_sign.
var meetingSignColumns = MeetingSignColumns{
	Id:             "id",
	MeetingUuid:    "meeting_uuid",
	SystemFileUuid: "system_file_uuid",
	SignStartTime:  "sign_start_time",
	Title:          "title",
	SignType:       "sign_type",
	Status:         "status",
	QrCode:         "qr_code",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewMeetingSignDao creates and returns a new DAO object for table data access.
func NewMeetingSignDao() *MeetingSignDao {
	return &MeetingSignDao{
		group:   "default",
		table:   "meeting_sign",
		columns: meetingSignColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingSignDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingSignDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingSignDao) Columns() MeetingSignColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingSignDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingSignDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingSignDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
