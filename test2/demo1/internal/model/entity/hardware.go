// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Hardware is the golang structure for table hardware.
type Hardware struct {
	Id            uint        `json:"id"            ` //
	CpuPercent    int         `json:"cpuPercent"    ` //
	MemoryPercent int         `json:"memoryPercent" ` //
	HarddiskSpace int         `json:"harddiskSpace" ` //
	TerminalMac   string      `json:"terminalMac"   ` //
	CreatedAt     *gtime.Time `json:"createdAt"     ` //
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` //
}
