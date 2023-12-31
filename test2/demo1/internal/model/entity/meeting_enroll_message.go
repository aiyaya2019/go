// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingEnrollMessage is the golang structure for table meeting_enroll_message.
type MeetingEnrollMessage struct {
	Id          uint        `json:"id"          ` //
	MeetingUuid uint        `json:"meetingUuid" ` //
	Contact     string      `json:"contact"     ` //
	Link        string      `json:"link"        ` //
	TimeType    int         `json:"timeType"    ` //
	StartTime   int         `json:"startTime"   ` //
	EndTime     int         `json:"endTime"     ` //
	QrcodePath  string      `json:"qrcodePath"  ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
