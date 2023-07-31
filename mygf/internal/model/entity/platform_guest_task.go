// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlatformGuestTask is the golang structure for table platform_guest_task.
type PlatformGuestTask struct {
	Id               int         `json:"id"               ` // 主键
	Name             string      `json:"name"             ` // 会议名字
	MeetingStatus    int         `json:"meetingStatus"    ` // 会议状态
	StartTime        *gtime.Time `json:"startTime"        ` // 开始时间
	EndTime          *gtime.Time `json:"endTime"          ` // 结束时间
	MeetingSlogan    string      `json:"meetingSlogan"    ` // 会议标语
	Nameplate        string      `json:"nameplate"        ` // 参会人铭牌
	IsSecrect        int         `json:"isSecrect"        ` // 会议设置(0为普通会议，1为保密会议，2开放会议)
	Description      string      `json:"description"      ` // 会议描述
	ModeratorName    string      `json:"moderatorName"    ` // 主席名字无关联
	CreateName       string      `json:"createName"       ` // 创建者名称
	MenuTab          int         `json:"menuTab"          ` // 选择的菜单数值
	AuthorityJson    string      `json:"authorityJson"    ` // 客户端自定义菜单
	HostIp           string      `json:"hostIp"           ` // 主会场ip
	HostPlatformUuid uint        `json:"hostPlatformUuid" ` // 主会场关联uuid，platform表的uuid
	HostPsIp         string      `json:"hostPsIp"         ` // 主会场主ps的ip
	HostPsPort       int         `json:"hostPsPort"       ` // 主会场主ps的port
	HostMeetingUuid  uint        `json:"hostMeetingUuid"  ` // 主会场关联会议
	Status           int         `json:"status"           ` // 状态 1:未加入2:已加入3:加入失败4:主会场结束5:主会场删除6:被移除7:已处理8:开会中9:已退出10:已拒绝11:副平台结束
	Operate          int         `json:"operate"          ` // 操作1add 2edit 3del
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
