// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MonitoringThreshold is the golang structure of table monitoring_threshold for DAO operations like Where/Data.
type MonitoringThreshold struct {
	g.Meta          `orm:"table:monitoring_threshold, do:true"`
	Id              interface{} //
	HarddiskPercent interface{} //
	MemoryPercent   interface{} //
	CpuPercent      interface{} //
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
