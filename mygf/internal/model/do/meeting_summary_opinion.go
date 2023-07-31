// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummaryOpinion is the golang structure of table meeting_summary_opinion for DAO operations like Where/Data.
type MeetingSummaryOpinion struct {
	g.Meta           `orm:"table:meeting_summary_opinion, do:true"`
	Id               interface{} //
	MeetingUuid      interface{} // 会议uuid，meeting表的uuid
	UserUuid         interface{} // 用户uuid，user表的uuid
	MeetingSummaryId interface{} // 纪要ID，meeting_summary表的id
	SendTime         *gtime.Time // 发送时间
	Content          interface{} // 纪要内容
	Page             interface{} // 在纪要文件中第几页提意见
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
