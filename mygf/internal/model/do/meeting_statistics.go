// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingStatistics is the golang structure of table meeting_statistics for DAO operations like Where/Data.
type MeetingStatistics struct {
	g.Meta           `orm:"table:meeting_statistics, do:true"`
	Id               interface{} // 主键
	MeetingUuid      interface{} // 关联会议uuid
	MeetingUserCount interface{} // 参会人数
	MeetingDuration  interface{} // 会议时长
	SavePaperCount   interface{} // 节约纸张数
	MeetingTime      *gtime.Time // 会议时间
	StartTime        *gtime.Time // 会议开始时间
	EndTime          *gtime.Time // 会议结束时间
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
