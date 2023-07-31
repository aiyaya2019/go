// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUser is the golang structure for table meeting_user.
type MeetingUser struct {
	Id                  uint        `json:"id"                  ` // 逻辑ID
	MeetingUuid         uint        `json:"meetingUuid"         ` // 会议uuid，meeting表的uuid
	UserUuid            uint        `json:"userUuid"            ` // 用户uuid，user表的uuid
	TerminalId          uint        `json:"terminalId"          ` // 终端ID;终端信息表中ID，terminal表的id
	IsBroadcast         uint        `json:"isBroadcast"         ` // 是否为广播:0=>不能广播，1=>广播
	Status              uint        `json:"status"              ` // 在会状态:0=>未在会，1=>在会
	IsSecretary         uint        `json:"isSecretary"         ` // 是否会务：0否，1是。当==0时，is_attendee=1
	Username            string      `json:"username"            ` // 用户名称
	Nameplate           string      `json:"nameplate"           ` // 个人铭牌/欢迎界面数据
	IsUpdateNameplate   int         `json:"isUpdateNameplate"   ` // 铭牌是否更新过：0否；1是
	IsUpdateWelcomePage int         `json:"isUpdateWelcomePage" ` // 欢迎界面是否更新过：0否；1是
	Mark                string      `json:"mark"                ` // 备注
	FreeLogin           uint        `json:"freeLogin"           ` // 是否免密登录,0否
	IsChairman          uint        `json:"isChairman"          ` // 是否为主席
	IsAttendee          int         `json:"isAttendee"          ` // 是否参会：0否；1是。当==0时，秘书和主持必须都=1
	Sort                int         `json:"sort"                ` // 排序
	Lifter              uint        `json:"lifter"              ` // 0不是升降器；1升降器，对应升降器id记录在terminal_id
	Deleted             int         `json:"deleted"             ` // 软删除，0正常，1已删除
	Unit                string      `json:"unit"                ` // 单位
	Position            string      `json:"position"            ` // 职务
	IsLocal             int         `json:"isLocal"             ` // 本地用户,1是
	Alias               string      `json:"alias"               ` // 铭牌显示名称
	AttendType          int         `json:"attendType"          ` // 1列席2出席3请假
	IsAutoQueue         uint        `json:"isAutoQueue"         ` // 是否自动排位
	IsSign              uint        `json:"isSign"              ` // 是否签到：0否；1是
	CreatedAt           *gtime.Time `json:"createdAt"           ` // 创建时间
	UpdatedAt           *gtime.Time `json:"updatedAt"           ` // 更新时间
}
