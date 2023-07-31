// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingDatumDao is the data access object for table meeting_datum.
type MeetingDatumDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns MeetingDatumColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingDatumColumns defines and stores column names for table meeting_datum.
type MeetingDatumColumns struct {
	Id               string // 逻辑ID
	Uuid             string // 议题uuid
	MeetingUuid      string // 会议uuid，meeting表的uuid
	UserUuid         string // 上传用户uuid，user表的uuid
	PlatformUuid     string // 平台uuid
	SystemFileUuid   string // 文件uuid
	Name             string // 议题名称
	Status           string // 状态:0=>未审议，1=>审议中，2=>已审议，3=>准备审议中
	Sort             string // 序号：用于文件排序
	Mark             string // 备注
	Reporter         string // 汇报人
	ReportTime       string // 汇报时间
	IsJoinVote       string // 是否加入投票，0否 1加入
	IsSecret         string // 0不保密1保密
	StartTime        string // 开始时间
	EndTime          string // 结束时间
	PushTime         string // 推送时间
	PushStatus       string // 0:未推送1:推送成功；1推送失败
	StartDatum       string // 0议题开启前显示,1议题开启后显示
	RoomSeatSchemeId string // 排位方案id，room_seat_scheme表的id
	ReporterField    string // 汇报人字段
	UserShow         string // 参会人是否显示（0隐藏，1显示）
	Attendee         string // 参会人按查看权限或自定义（0查看权限，1自定义）
	UserDefined      string // 用户自定义字段
	AttendeeField    string // 参会人字段
	ReporterShow     string // 汇报人是否显示（0隐藏，1显示）
	UnitShow         string // 汇报单位是否显示（0隐藏，1显示）
	UnitField        string // 汇报单位字段
	Unit             string // 汇报单位
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// meetingDatumColumns holds the columns for table meeting_datum.
var meetingDatumColumns = MeetingDatumColumns{
	Id:               "id",
	Uuid:             "uuid",
	MeetingUuid:      "meeting_uuid",
	UserUuid:         "user_uuid",
	PlatformUuid:     "platform_uuid",
	SystemFileUuid:   "system_file_uuid",
	Name:             "name",
	Status:           "status",
	Sort:             "sort",
	Mark:             "mark",
	Reporter:         "reporter",
	ReportTime:       "report_time",
	IsJoinVote:       "is_join_vote",
	IsSecret:         "is_secret",
	StartTime:        "start_time",
	EndTime:          "end_time",
	PushTime:         "push_time",
	PushStatus:       "push_status",
	StartDatum:       "start_datum",
	RoomSeatSchemeId: "room_seat_scheme_id",
	ReporterField:    "reporter_field",
	UserShow:         "user_show",
	Attendee:         "attendee",
	UserDefined:      "user_defined",
	AttendeeField:    "attendee_field",
	ReporterShow:     "reporter_show",
	UnitShow:         "unit_show",
	UnitField:        "unit_field",
	Unit:             "unit",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewMeetingDatumDao creates and returns a new DAO object for table data access.
func NewMeetingDatumDao() *MeetingDatumDao {
	return &MeetingDatumDao{
		group:   "default",
		table:   "meeting_datum",
		columns: meetingDatumColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingDatumDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingDatumDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingDatumDao) Columns() MeetingDatumColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingDatumDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingDatumDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingDatumDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
