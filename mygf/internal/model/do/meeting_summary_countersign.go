// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummaryCountersign is the golang structure of table meeting_summary_countersign for DAO operations like Where/Data.
type MeetingSummaryCountersign struct {
	g.Meta             `orm:"table:meeting_summary_countersign, do:true"`
	MeetingUuid        interface{} // 会议uuid，meeting表的uuid
	MeetingSummaryId   interface{} // 会议纪要ID，meeting_summary表的id
	UserSystemFileUuid interface{} // 会签文件uuid，user_file的system_file_uuid
	UserUuid           interface{} // 用户uuid，user表的uuid
	Status             interface{} // 会签状态
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
}
