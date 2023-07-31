// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskQueue is the golang structure for table task_queue.
type TaskQueue struct {
	Id          uint        `json:"id"          ` // 逻辑ID
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	Inst        uint        `json:"inst"        ` // 任务指令，协议定的指令号
	ServerId    uint        `json:"serverId"    ` // 服务器ID，info_server表的id
	InsertTime  *gtime.Time `json:"insertTime"  ` // 插入时间
	ExecTime    *gtime.Time `json:"execTime"    ` // 执行时间
	Sort        uint        `json:"sort"        ` // 权重
	Content     string      `json:"content"     ` // json
	Status      uint        `json:"status"      ` // 状态： 0=》等待，1=》进行中，2=》成功，3=》失败，4=》服务器离线
	Mark        string      `json:"mark"        ` // 备注
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
