// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Hardware is the golang structure of table hardware for DAO operations like Where/Data.
type Hardware struct {
	g.Meta        `orm:"table:hardware, do:true"`
	Id            interface{} //
	CpuPercent    interface{} // cpu使用率
	MemoryPercent interface{} // 内存使用率
	HarddiskSpace interface{} // 硬盘剩余空间
	TerminalMac   interface{} // 终端mac，默认1为8300服务器
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
