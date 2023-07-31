// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingExtra is the golang structure of table meeting_extra for DAO operations like Where/Data.
type MeetingExtra struct {
	g.Meta          `orm:"table:meeting_extra, do:true"`
	MeetingUuid     interface{} // 会议uuid，meeting表的uuid
	MenuTab         interface{} // 选择的菜单数值
	NameplateCustom interface{} // 铭牌的自定义文本，限制50个汉字
	WelcomeCustom   interface{} // 欢迎界面的自定义文本，限制50个汉字
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
