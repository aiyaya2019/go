// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure for table operation_log.
type OperationLog struct {
	Id            int         `json:"id"            ` // 逻辑ID
	UserUuid      uint        `json:"userUuid"      ` // 操作用户uuid，user表的uuid
	Module        string      `json:"module"        ` // 模块
	Operation     string      `json:"operation"     ` // 操作
	OperationType string      `json:"operationType" ` // 操作类型
	Detail        string      `json:"detail"        ` // 操作详细信息
	CreatedAt     *gtime.Time `json:"createdAt"     ` // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` // 更新时间
}
