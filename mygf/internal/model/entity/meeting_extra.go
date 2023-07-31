// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingExtra is the golang structure for table meeting_extra.
type MeetingExtra struct {
	MeetingUuid     uint        `json:"meetingUuid"     ` // 会议uuid，meeting表的uuid
	MenuTab         int         `json:"menuTab"         ` // 选择的菜单数值
	NameplateCustom string      `json:"nameplateCustom" ` // 铭牌的自定义文本，限制50个汉字
	WelcomeCustom   string      `json:"welcomeCustom"   ` // 欢迎界面的自定义文本，限制50个汉字
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
}
