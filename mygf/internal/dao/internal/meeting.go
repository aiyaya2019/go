// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingDao is the data access object for table meeting.
type MeetingDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns MeetingColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingColumns defines and stores column names for table meeting.
type MeetingColumns struct {
	Id               string // 会议ID
	PlatformUuid     string // 平台uuid，platform表的uuid
	Uuid             string // 会议分布式uuid
	UserUuid         string // 创建者用户uuid，user表的uuid
	Name             string // 会议名称
	Type             string // 类型:0=>不要登录，1=>要登录
	Moderator        string // 会议主持人，user表的uuid
	Secretary        string // 会议秘书
	StartTime        string // 开始时间
	EndTime          string // 结束时间
	RealStartTime    string // 真正开始时间
	RealEndTime      string // 真正结束时间
	MasterRoomId     string // 主会议室id,room表的id
	RoomListId       string // 会议室ID;多个会议室ID用’,’号隔开
	MeetingSlogan    string // 会议标语
	Nameplate        string // 参会人铭牌
	Status           string // 状态:0未开始，1开会中，2结束,3模板会议,4暂停
	IsSecrect        string // 会议设置(0为普通会议，1为保密会议，2开放会议)
	Description      string // 会议描述
	AllUser          string // 0手选,1全选
	ModeratorName    string // 主席名字无关联
	MeetingType      string // 会议类型,1常规2本地3异地
	IsHost           string // 1.主会场2.客会场
	HostMeetingUuid  string // 主会场关联uuid，meeting表的uuid
	HostPlatformUuid string // 主会场平台uuid，platform表的uuid
	HostIp           string // 主会场ip
	HostPsIp         string // 主会场主ps的ip
	HostPsPort       string // 主会场主ps的port
	MenuData         string // 会议菜单信息
	MenuTab          string // 菜单模式：1卡片模式；2简洁模式;3导航模式功能菜单；4经典模式功能菜单
	NameplateCustom  string // 铭牌的自定义文本，限制50个汉字
	WelcomeCustom    string // 欢迎界面的自定义文本，限制50个汉字
	CustomBgImg      string // 自定义背景图
	Deleted          string // 删除：0未删除；1删除到回收站；2删除
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// meetingColumns holds the columns for table meeting.
var meetingColumns = MeetingColumns{
	Id:               "id",
	PlatformUuid:     "platform_uuid",
	Uuid:             "uuid",
	UserUuid:         "user_uuid",
	Name:             "name",
	Type:             "type",
	Moderator:        "moderator",
	Secretary:        "secretary",
	StartTime:        "start_time",
	EndTime:          "end_time",
	RealStartTime:    "real_start_time",
	RealEndTime:      "real_end_time",
	MasterRoomId:     "master_room_id",
	RoomListId:       "room_list_id",
	MeetingSlogan:    "meeting_slogan",
	Nameplate:        "nameplate",
	Status:           "status",
	IsSecrect:        "is_secrect",
	Description:      "description",
	AllUser:          "all_user",
	ModeratorName:    "moderator_name",
	MeetingType:      "meeting_type",
	IsHost:           "is_host",
	HostMeetingUuid:  "host_meeting_uuid",
	HostPlatformUuid: "host_platform_uuid",
	HostIp:           "host_ip",
	HostPsIp:         "host_ps_ip",
	HostPsPort:       "host_ps_port",
	MenuData:         "menu_data",
	MenuTab:          "menu_tab",
	NameplateCustom:  "nameplate_custom",
	WelcomeCustom:    "welcome_custom",
	CustomBgImg:      "custom_bg_img",
	Deleted:          "deleted",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewMeetingDao creates and returns a new DAO object for table data access.
func NewMeetingDao() *MeetingDao {
	return &MeetingDao{
		group:   "default",
		table:   "meeting",
		columns: meetingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingDao) Columns() MeetingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
