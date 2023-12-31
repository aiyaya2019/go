// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomAppointment is the golang structure for table room_appointment.
type RoomAppointment struct {
	Id          int         `json:"id"          ` //
	MeetingUuid uint        `json:"meetingUuid" ` //
	UserUuid    uint        `json:"userUuid"    ` //
	RoomId      int         `json:"roomId"      ` //
	Name        string      `json:"name"        ` //
	StartTime   *gtime.Time `json:"startTime"   ` //
	EndTime     *gtime.Time `json:"endTime"     ` //
	Status      int         `json:"status"      ` //
	Note        string      `json:"note"        ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
