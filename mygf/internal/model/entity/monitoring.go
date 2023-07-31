// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Monitoring is the golang structure for table monitoring.
type Monitoring struct {
	Id               uint        `json:"id"               ` //
	HarddiskUse      int         `json:"harddiskUse"      ` // 硬盘使用空间
	Harddisk         int         `json:"harddisk"         ` // 硬盘
	MemoryUse        int         `json:"memoryUse"        ` // 内存使用
	Memory           int         `json:"memory"           ` // 内存
	CpuPercent       int         `json:"cpuPercent"       ` // cpu使用率
	NetworkUpspeed   string      `json:"networkUpspeed"   ` // 上行网速
	NetworkDownspeed string      `json:"networkDownspeed" ` // 下行网速
	TerminalMac      string      `json:"terminalMac"      ` // 终端mac，默认1为8300服务器
	LastViewTime     *gtime.Time `json:"lastViewTime"     ` // 最后一次浏览通知时间
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
