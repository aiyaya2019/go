// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingDatumUser is the golang structure for table meeting_datum_user.
type MeetingDatumUser struct {
	MeetingUuid      uint        `json:"meetingUuid"      ` //
	UserUuid         uint        `json:"userUuid"         ` //
	MeetingDatumUuid uint        `json:"meetingDatumUuid" ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
}