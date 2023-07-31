// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TerminalRepair is the golang structure for table terminal_repair.
type TerminalRepair struct {
	Id          int         `json:"id"          ` // 主键
	TerminalIds string      `json:"terminalIds" ` // 终端ids，terminal表的id
	FileName    string      `json:"fileName"    ` // 文件名称
	Command     string      `json:"command"     ` // 命令/客户端存放路径
	Path        string      `json:"path"        ` // 上传文件路径
	Type        int         `json:"type"        ` // 类型1->命令,2->文件
	Status      int         `json:"status"      ` // 1=>进行中,2=>成功,3=>失败
	CreateTime  *gtime.Time `json:"createTime"  ` // 创建时间
	Mark        string      `json:"mark"        ` // 备注
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
