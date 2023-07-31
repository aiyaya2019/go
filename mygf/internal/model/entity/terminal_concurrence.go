// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TerminalConcurrence is the golang structure for table terminal_concurrence.
type TerminalConcurrence struct {
	Id          int         `json:"id"          ` // 主键
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	TotalUser   int         `json:"totalUser"   ` // 单节点客户端并发数
	ClientIds   string      `json:"clientIds"   ` // 勾选的客户端id逗号隔开，type=8
	Usernames   string      `json:"usernames"   ` // 随机生成用户名erupt_test1
	UserMax     int         `json:"userMax"     ` // 最大用户值
	Type        int         `json:"type"        ` // 操作1->添加3->删除
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
