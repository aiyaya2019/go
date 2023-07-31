// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingSignInfoDao is the data access object for table meeting_sign_info.
type MeetingSignInfoDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns MeetingSignInfoColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingSignInfoColumns defines and stores column names for table meeting_sign_info.
type MeetingSignInfoColumns struct {
	Id             string // 逻辑ID
	MeetingUuid    string // 会议uuid，meeting表的uuid
	UserUuid       string // 用户uuid，user表的uuid
	SystemFileUuid string // 签到文件uuid，system_file表的uuid
	MeetingSignId  string // 签到标记ID，meeting_sign表的id
	Status         string // 签到状态：0未签到；1已签到
	SignTime       string // 签到时间
	Assist         string // 是否协助签到，1:是
	DeviceType     string // 设备类型：1无纸化；2电子桌牌
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// meetingSignInfoColumns holds the columns for table meeting_sign_info.
var meetingSignInfoColumns = MeetingSignInfoColumns{
	Id:             "id",
	MeetingUuid:    "meeting_uuid",
	UserUuid:       "user_uuid",
	SystemFileUuid: "system_file_uuid",
	MeetingSignId:  "meeting_sign_id",
	Status:         "status",
	SignTime:       "sign_time",
	Assist:         "assist",
	DeviceType:     "device_type",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewMeetingSignInfoDao creates and returns a new DAO object for table data access.
func NewMeetingSignInfoDao() *MeetingSignInfoDao {
	return &MeetingSignInfoDao{
		group:   "default",
		table:   "meeting_sign_info",
		columns: meetingSignInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingSignInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingSignInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingSignInfoDao) Columns() MeetingSignInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingSignInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingSignInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingSignInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
