// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Register is the golang structure for table register.
type Register struct {
	Id        uint        `json:"id"        ` // 逻辑ID
	Content   string      `json:"content"   ` // 注册码
	Status    uint        `json:"status"    ` // 状态
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
