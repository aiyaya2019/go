// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Meeting is the golang structure for table meeting.
type Meeting struct {
	Id               uint        `json:"id"               ` // 会议ID
	PlatformUuid     uint        `json:"platformUuid"     ` // 平台uuid，platform表的uuid
	Uuid             uint        `json:"uuid"             ` // 会议分布式uuid
	UserUuid         uint        `json:"userUuid"         ` // 创建者用户uuid，user表的uuid
	Name             string      `json:"name"             ` // 会议名称
	Type             uint        `json:"type"             ` // 类型:0=>不要登录，1=>要登录
	Moderator        uint        `json:"moderator"        ` // 会议主持人，user表的uuid
	Secretary        uint        `json:"secretary"        ` // 会议秘书
	StartTime        *gtime.Time `json:"startTime"        ` // 开始时间
	EndTime          *gtime.Time `json:"endTime"          ` // 结束时间
	RealStartTime    *gtime.Time `json:"realStartTime"    ` // 真正开始时间
	RealEndTime      *gtime.Time `json:"realEndTime"      ` // 真正结束时间
	MasterRoomId     int         `json:"masterRoomId"     ` // 主会议室id,room表的id
	RoomListId       string      `json:"roomListId"       ` // 会议室ID;多个会议室ID用’,’号隔开
	MeetingSlogan    string      `json:"meetingSlogan"    ` // 会议标语
	Nameplate        string      `json:"nameplate"        ` // 参会人铭牌
	Status           uint        `json:"status"           ` // 状态:0未开始，1开会中，2结束,3模板会议,4暂停
	IsSecrect        uint        `json:"isSecrect"        ` // 会议设置(0为普通会议，1为保密会议，2开放会议)
	Description      string      `json:"description"      ` // 会议描述
	AllUser          int         `json:"allUser"          ` // 0手选,1全选
	ModeratorName    string      `json:"moderatorName"    ` // 主席名字无关联
	MeetingType      int         `json:"meetingType"      ` // 会议类型,1常规2本地3异地
	IsHost           int         `json:"isHost"           ` // 1.主会场2.客会场
	HostMeetingUuid  uint        `json:"hostMeetingUuid"  ` // 主会场关联uuid，meeting表的uuid
	HostPlatformUuid uint        `json:"hostPlatformUuid" ` // 主会场平台uuid，platform表的uuid
	HostIp           string      `json:"hostIp"           ` // 主会场ip
	HostPsIp         string      `json:"hostPsIp"         ` // 主会场主ps的ip
	HostPsPort       int         `json:"hostPsPort"       ` // 主会场主ps的port
	MenuData         string      `json:"menuData"         ` // 会议菜单信息
	MenuTab          int         `json:"menuTab"          ` // 菜单模式：1卡片模式；2简洁模式;3导航模式功能菜单；4经典模式功能菜单
	NameplateCustom  string      `json:"nameplateCustom"  ` // 铭牌的自定义文本，限制50个汉字
	WelcomeCustom    string      `json:"welcomeCustom"    ` // 欢迎界面的自定义文本，限制50个汉字
	CustomBgImg      string      `json:"customBgImg"      ` // 自定义背景图
	Deleted          int         `json:"deleted"          ` // 删除：0未删除；1删除到回收站；2删除
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
