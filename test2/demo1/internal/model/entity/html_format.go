// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HtmlFormat is the golang structure for table html_format.
type HtmlFormat struct {
	Id          uint        `json:"id"          ` //
	MeetingUuid uint        `json:"meetingUuid" ` //
	MultiUuid   uint        `json:"multiUuid"   ` //
	Url         string      `json:"url"         ` //
	Path        string      `json:"path"        ` //
	Progress    string      `json:"progress"    ` //
	Status      uint        `json:"status"      ` //
	Type        uint        `json:"type"        ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}