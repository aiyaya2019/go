// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomAppointment is the golang structure for table room_appointment.
type RoomAppointment struct {
	Id          int         `json:"id"          ` // 逻辑ID
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	UserUuid    uint        `json:"userUuid"    ` // 预约用户的uuid，user表的uuid
	RoomId      int         `json:"roomId"      ` // 会议室ID,room表id
	Name        string      `json:"name"        ` // 预约人员的名字
	StartTime   *gtime.Time `json:"startTime"   ` // 预约的开始时间
	EndTime     *gtime.Time `json:"endTime"     ` // 预约的结束时间
	Status      int         `json:"status"      ` // 状态，1正常，0停用
	Note        string      `json:"note"        ` // 备注
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
