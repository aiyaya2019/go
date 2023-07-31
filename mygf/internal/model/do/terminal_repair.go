// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TerminalRepair is the golang structure of table terminal_repair for DAO operations like Where/Data.
type TerminalRepair struct {
	g.Meta      `orm:"table:terminal_repair, do:true"`
	Id          interface{} // 主键
	TerminalIds interface{} // 终端ids，terminal表的id
	FileName    interface{} // 文件名称
	Command     interface{} // 命令/客户端存放路径
	Path        interface{} // 上传文件路径
	Type        interface{} // 类型1->命令,2->文件
	Status      interface{} // 1=>进行中,2=>成功,3=>失败
	CreateTime  *gtime.Time // 创建时间
	Mark        interface{} // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
