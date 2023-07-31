// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummaryOpinion is the golang structure for table meeting_summary_opinion.
type MeetingSummaryOpinion struct {
	Id               uint        `json:"id"               ` //
	MeetingUuid      uint        `json:"meetingUuid"      ` //
	UserUuid         uint        `json:"userUuid"         ` //
	MeetingSummaryId uint        `json:"meetingSummaryId" ` //
	SendTime         *gtime.Time `json:"sendTime"         ` //
	Content          string      `json:"content"          ` //
	Page             int         `json:"page"             ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
}
