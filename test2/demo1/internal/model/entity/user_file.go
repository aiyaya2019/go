// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFile is the golang structure for table user_file.
type UserFile struct {
	MeetingUuid    uint        `json:"meetingUuid"    ` //
	UserUuid       uint        `json:"userUuid"       ` //
	SystemFileUuid uint        `json:"systemFileUuid" ` //
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` //
}