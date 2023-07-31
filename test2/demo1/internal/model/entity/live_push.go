// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LivePush is the golang structure for table live_push.
type LivePush struct {
	Id             uint        `json:"id"             ` //
	SystemFileUuid uint        `json:"systemFileUuid" ` //
	Pid            int         `json:"pid"            ` //
	Progress       string      `json:"progress"       ` //
	Status         uint        `json:"status"         ` //
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` //
}
