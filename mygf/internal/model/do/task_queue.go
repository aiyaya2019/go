// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskQueue is the golang structure of table task_queue for DAO operations like Where/Data.
type TaskQueue struct {
	g.Meta      `orm:"table:task_queue, do:true"`
	Id          interface{} // 逻辑ID
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	Inst        interface{} // 任务指令，协议定的指令号
	ServerId    interface{} // 服务器ID，info_server表的id
	InsertTime  *gtime.Time // 插入时间
	ExecTime    *gtime.Time // 执行时间
	Sort        interface{} // 权重
	Content     interface{} // json
	Status      interface{} // 状态： 0=》等待，1=》进行中，2=》成功，3=》失败，4=》服务器离线
	Mark        interface{} // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
