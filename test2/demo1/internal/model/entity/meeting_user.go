// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUser is the golang structure for table meeting_user.
type MeetingUser struct {
	Id                  uint        `json:"id"                  ` //
	MeetingUuid         uint        `json:"meetingUuid"         ` //
	UserUuid            uint        `json:"userUuid"            ` //
	TerminalId          uint        `json:"terminalId"          ` //
	IsBroadcast         uint        `json:"isBroadcast"         ` //
	Status              uint        `json:"status"              ` //
	IsSecretary         uint        `json:"isSecretary"         ` //
	Username            string      `json:"username"            ` //
	Nameplate           string      `json:"nameplate"           ` //
	IsUpdateNameplate   int         `json:"isUpdateNameplate"   ` //
	IsUpdateWelcomePage int         `json:"isUpdateWelcomePage" ` //
	Mark                string      `json:"mark"                ` //
	FreeLogin           uint        `json:"freeLogin"           ` //
	IsChairman          uint        `json:"isChairman"          ` //
	IsAttendee          int         `json:"isAttendee"          ` //
	Sort                int         `json:"sort"                ` //
	Lifter              uint        `json:"lifter"              ` //
	Deleted             int         `json:"deleted"             ` //
	Unit                string      `json:"unit"                ` //
	Position            string      `json:"position"            ` //
	IsLocal             int         `json:"isLocal"             ` //
	Alias               string      `json:"alias"               ` //
	AttendType          int         `json:"attendType"          ` //
	IsAutoQueue         uint        `json:"isAutoQueue"         ` //
	IsSign              uint        `json:"isSign"              ` //
	CreatedAt           *gtime.Time `json:"createdAt"           ` //
	UpdatedAt           *gtime.Time `json:"updatedAt"           ` //
}
