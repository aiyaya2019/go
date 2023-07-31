// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSignInfo is the golang structure for table meeting_sign_info.
type MeetingSignInfo struct {
	Id             uint        `json:"id"             ` //
	MeetingUuid    uint        `json:"meetingUuid"    ` //
	UserUuid       uint        `json:"userUuid"       ` //
	SystemFileUuid uint        `json:"systemFileUuid" ` //
	MeetingSignId  uint        `json:"meetingSignId"  ` //
	Status         int         `json:"status"         ` //
	SignTime       *gtime.Time `json:"signTime"       ` //
	Assist         int         `json:"assist"         ` //
	DeviceType     int         `json:"deviceType"     ` //
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` //
}
