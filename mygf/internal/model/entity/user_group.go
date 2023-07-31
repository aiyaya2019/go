// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGroup is the golang structure for table user_group.
type UserGroup struct {
	Id        uint        `json:"id"        ` // 逻辑ID
	Name      string      `json:"name"      ` // 名称
	Mark      string      `json:"mark"      ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
