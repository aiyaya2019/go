// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskQueue is the golang structure for table task_queue.
type TaskQueue struct {
	Id          uint        `json:"id"          ` //
	MeetingUuid uint        `json:"meetingUuid" ` //
	Inst        uint        `json:"inst"        ` //
	ServerId    uint        `json:"serverId"    ` //
	InsertTime  *gtime.Time `json:"insertTime"  ` //
	ExecTime    *gtime.Time `json:"execTime"    ` //
	Sort        uint        `json:"sort"        ` //
	Content     string      `json:"content"     ` //
	Status      uint        `json:"status"      ` //
	Mark        string      `json:"mark"        ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
