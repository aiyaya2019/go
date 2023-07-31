// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingEnrollMessage is the golang structure for table meeting_enroll_message.
type MeetingEnrollMessage struct {
	Id          uint        `json:"id"          ` // 逻辑id
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	Contact     string      `json:"contact"     ` // 报名联系
	Link        string      `json:"link"        ` // 报名链接
	TimeType    int         `json:"timeType"    ` // 0:开始到结束时间内报名;1:会议状态开始前可以报名
	StartTime   int         `json:"startTime"   ` // 开始时间
	EndTime     int         `json:"endTime"     ` // 结束时间
	QrcodePath  string      `json:"qrcodePath"  ` // 二维码保存路径
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
