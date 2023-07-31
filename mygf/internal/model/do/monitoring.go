// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
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
	HarddiskUse      interface{} // 硬盘使用空间
	Harddisk         interface{} // 硬盘
	MemoryUse        interface{} // 内存使用
	Memory           interface{} // 内存
	CpuPercent       interface{} // cpu使用率
	NetworkUpspeed   interface{} // 上行网速
	NetworkDownspeed interface{} // 下行网速
	TerminalMac      interface{} // 终端mac，默认1为8300服务器
	LastViewTime     *gtime.Time // 最后一次浏览通知时间
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
