// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingRecordContent is the golang structure of table meeting_record_content for DAO operations like Where/Data.
type MeetingRecordContent struct {
	g.Meta      `orm:"table:meeting_record_content, do:true"`
	Id          interface{} //
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	TerminalId  interface{} // 终端ID，terminal表的id
	HtmlContent interface{} // HTML会议记录内容
	WordContent interface{} // Word会议记录内容
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
