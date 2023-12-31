// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Monitoring is the golang structure of table monitoring for DAO operations like Where/Data.
type Monitoring struct {
	g.Meta           `orm:"table:monitoring, do:true"`
	Id               interface{} //
	HarddiskUse      interface{} //
	Harddisk         interface{} //
	MemoryUse        interface{} //
	Memory           interface{} //
	CpuPercent       interface{} //
	NetworkUpspeed   interface{} //
	NetworkDownspeed interface{} //
	TerminalMac      interface{} //
	LastViewTime     *gtime.Time //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}
