// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Hardware is the golang structure for table hardware.
type Hardware struct {
	Id            uint        `json:"id"            ` //
	CpuPercent    int         `json:"cpuPercent"    ` // cpu使用率
	MemoryPercent int         `json:"memoryPercent" ` // 内存使用率
	HarddiskSpace int         `json:"harddiskSpace" ` // 硬盘剩余空间
	TerminalMac   string      `json:"terminalMac"   ` // 终端mac，默认1为8300服务器
	CreatedAt     *gtime.Time `json:"createdAt"     ` // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` // 更新时间
}
