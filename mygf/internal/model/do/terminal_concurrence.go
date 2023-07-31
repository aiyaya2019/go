// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TerminalConcurrence is the golang structure of table terminal_concurrence for DAO operations like Where/Data.
type TerminalConcurrence struct {
	g.Meta      `orm:"table:terminal_concurrence, do:true"`
	Id          interface{} // 主键
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	TotalUser   interface{} // 单节点客户端并发数
	ClientIds   interface{} // 勾选的客户端id逗号隔开，type=8
	Usernames   interface{} // 随机生成用户名erupt_test1
	UserMax     interface{} // 最大用户值
	Type        interface{} // 操作1->添加3->删除
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
