// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TerminalRepair is the golang structure for table terminal_repair.
type TerminalRepair struct {
	Id          int         `json:"id"          ` //
	TerminalIds string      `json:"terminalIds" ` //
	FileName    string      `json:"fileName"    ` //
	Command     string      `json:"command"     ` //
	Path        string      `json:"path"        ` //
	Type        int         `json:"type"        ` //
	Status      int         `json:"status"      ` //
	CreateTime  *gtime.Time `json:"createTime"  ` //
	Mark        string      `json:"mark"        ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
