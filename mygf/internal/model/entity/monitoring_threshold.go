// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MonitoringThreshold is the golang structure for table monitoring_threshold.
type MonitoringThreshold struct {
	Id              uint        `json:"id"              ` // 主键
	HarddiskPercent int         `json:"harddiskPercent" ` // 磁盘使用率
	MemoryPercent   int         `json:"memoryPercent"   ` // 内存使用率
	CpuPercent      int         `json:"cpuPercent"      ` // CPU百分比
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
}
