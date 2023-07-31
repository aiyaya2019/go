// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummaryThird is the golang structure of table meeting_summary_third for DAO operations like Where/Data.
type MeetingSummaryThird struct {
	g.Meta           `orm:"table:meeting_summary_third, do:true"`
	MeetingSummaryId interface{} // 会议纪要id，meeting_summary的id
	PhoneticTranId   interface{} // 语音转写那边的id，作对应关系
	Creator          interface{} // 创建者
	Modifier         interface{} // 修改者
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
