// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingDatum is the golang structure for table meeting_datum.
type MeetingDatum struct {
	Id               uint        `json:"id"               ` // 逻辑ID
	Uuid             uint        `json:"uuid"             ` // 议题uuid
	MeetingUuid      uint        `json:"meetingUuid"      ` // 会议uuid，meeting表的uuid
	UserUuid         uint        `json:"userUuid"         ` // 上传用户uuid，user表的uuid
	PlatformUuid     uint        `json:"platformUuid"     ` // 平台uuid
	SystemFileUuid   uint        `json:"systemFileUuid"   ` // 文件uuid
	Name             string      `json:"name"             ` // 议题名称
	Status           uint        `json:"status"           ` // 状态:0=>未审议，1=>审议中，2=>已审议，3=>准备审议中
	Sort             int         `json:"sort"             ` // 序号：用于文件排序
	Mark             string      `json:"mark"             ` // 备注
	Reporter         string      `json:"reporter"         ` // 汇报人
	ReportTime       string      `json:"reportTime"       ` // 汇报时间
	IsJoinVote       uint        `json:"isJoinVote"       ` // 是否加入投票，0否 1加入
	IsSecret         int         `json:"isSecret"         ` // 0不保密1保密
	StartTime        *gtime.Time `json:"startTime"        ` // 开始时间
	EndTime          *gtime.Time `json:"endTime"          ` // 结束时间
	PushTime         int         `json:"pushTime"         ` // 推送时间
	PushStatus       int         `json:"pushStatus"       ` // 0:未推送1:推送成功；1推送失败
	StartDatum       int         `json:"startDatum"       ` // 0议题开启前显示,1议题开启后显示
	RoomSeatSchemeId int         `json:"roomSeatSchemeId" ` // 排位方案id，room_seat_scheme表的id
	ReporterField    string      `json:"reporterField"    ` // 汇报人字段
	UserShow         int         `json:"userShow"         ` // 参会人是否显示（0隐藏，1显示）
	Attendee         int         `json:"attendee"         ` // 参会人按查看权限或自定义（0查看权限，1自定义）
	UserDefined      string      `json:"userDefined"      ` // 用户自定义字段
	AttendeeField    string      `json:"attendeeField"    ` // 参会人字段
	ReporterShow     int         `json:"reporterShow"     ` // 汇报人是否显示（0隐藏，1显示）
	UnitShow         int         `json:"unitShow"         ` // 汇报单位是否显示（0隐藏，1显示）
	UnitField        string      `json:"unitField"        ` // 汇报单位字段
	Unit             string      `json:"unit"             ` // 汇报单位
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
