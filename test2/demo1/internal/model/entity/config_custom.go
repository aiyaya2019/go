// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigCustom is the golang structure for table config_custom.
type ConfigCustom struct {
	Id             int         `json:"id"             ` //
	SysName        string      `json:"sysName"        ` //
	Logo           string      `json:"logo"           ` //
	Background     string      `json:"background"     ` //
	DefaultSetting int         `json:"defaultSetting" ` //
	Type           int         `json:"type"           ` //
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` //
}
