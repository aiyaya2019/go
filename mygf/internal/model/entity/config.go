// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure for table config.
type Config struct {
	Id        int         `json:"id"        ` // 配置id
	Key       string      `json:"key"       ` // 配置键名标识
	Value     string      `json:"value"     ` // 配置内容
	Default   string      `json:"default"   ` // 默认配置内容
	Mark      string      `json:"mark"      ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
