// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingUserDao is the data access object for table meeting_user.
type MeetingUserDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns MeetingUserColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingUserColumns defines and stores column names for table meeting_user.
type MeetingUserColumns struct {
	Id                  string // 逻辑ID
	MeetingUuid         string // 会议uuid，meeting表的uuid
	UserUuid            string // 用户uuid，user表的uuid
	TerminalId          string // 终端ID;终端信息表中ID，terminal表的id
	IsBroadcast         string // 是否为广播:0=>不能广播，1=>广播
	Status              string // 在会状态:0=>未在会，1=>在会
	IsSecretary         string // 是否会务：0否，1是。当==0时，is_attendee=1
	Username            string // 用户名称
	Nameplate           string // 个人铭牌/欢迎界面数据
	IsUpdateNameplate   string // 铭牌是否更新过：0否；1是
	IsUpdateWelcomePage string // 欢迎界面是否更新过：0否；1是
	Mark                string // 备注
	FreeLogin           string // 是否免密登录,0否
	IsChairman          string // 是否为主席
	IsAttendee          string // 是否参会：0否；1是。当==0时，秘书和主持必须都=1
	Sort                string // 排序
	Lifter              string // 0不是升降器；1升降器，对应升降器id记录在terminal_id
	Deleted             string // 软删除，0正常，1已删除
	Unit                string // 单位
	Position            string // 职务
	IsLocal             string // 本地用户,1是
	Alias               string // 铭牌显示名称
	AttendType          string // 1列席2出席3请假
	IsAutoQueue         string // 是否自动排位
	IsSign              string // 是否签到：0否；1是
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
}

// meetingUserColumns holds the columns for table meeting_user.
var meetingUserColumns = MeetingUserColumns{
	Id:                  "id",
	MeetingUuid:         "meeting_uuid",
	UserUuid:            "user_uuid",
	TerminalId:          "terminal_id",
	IsBroadcast:         "is_broadcast",
	Status:              "status",
	IsSecretary:         "is_secretary",
	Username:            "username",
	Nameplate:           "nameplate",
	IsUpdateNameplate:   "is_update_nameplate",
	IsUpdateWelcomePage: "is_update_welcome_page",
	Mark:                "mark",
	FreeLogin:           "free_login",
	IsChairman:          "is_chairman",
	IsAttendee:          "is_attendee",
	Sort:                "sort",
	Lifter:              "lifter",
	Deleted:             "deleted",
	Unit:                "unit",
	Position:            "position",
	IsLocal:             "is_local",
	Alias:               "alias",
	AttendType:          "attend_type",
	IsAutoQueue:         "is_auto_queue",
	IsSign:              "is_sign",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
}

// NewMeetingUserDao creates and returns a new DAO object for table data access.
func NewMeetingUserDao() *MeetingUserDao {
	return &MeetingUserDao{
		group:   "default",
		table:   "meeting_user",
		columns: meetingUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingUserDao) Columns() MeetingUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
