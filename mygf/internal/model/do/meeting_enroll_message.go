// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingEnrollMessage is the golang structure of table meeting_enroll_message for DAO operations like Where/Data.
type MeetingEnrollMessage struct {
	g.Meta      `orm:"table:meeting_enroll_message, do:true"`
	Id          interface{} // 逻辑id
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	Contact     interface{} // 报名联系
	Link        interface{} // 报名链接
	TimeType    interface{} // 0:开始到结束时间内报名;1:会议状态开始前可以报名
	StartTime   interface{} // 开始时间
	EndTime     interface{} // 结束时间
	QrcodePath  interface{} // 二维码保存路径
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
