// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUserRights is the golang structure for table meeting_user_rights.
type MeetingUserRights struct {
	Id          int         `json:"id"          ` //
	MeetingUuid uint        `json:"meetingUuid" ` //
	UserUuid    uint        `json:"userUuid"    ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
