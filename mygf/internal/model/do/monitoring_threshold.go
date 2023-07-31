// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MonitoringThreshold is the golang structure of table monitoring_threshold for DAO operations like Where/Data.
type MonitoringThreshold struct {
	g.Meta          `orm:"table:monitoring_threshold, do:true"`
	Id              interface{} // 主键
	HarddiskPercent interface{} // 磁盘使用率
	MemoryPercent   interface{} // 内存使用率
	CpuPercent      interface{} // CPU百分比
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
