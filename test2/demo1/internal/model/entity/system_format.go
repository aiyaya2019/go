// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemFormat is the golang structure for table system_format.
type SystemFormat struct {
	Id           int         `json:"id"           ` //
	Type         int         `json:"type"         ` //
	UploadSize   string      `json:"uploadSize"   ` //
	Size         int         `json:"size"         ` //
	SystemFormat string      `json:"systemFormat" ` //
	DefineFormat string      `json:"defineFormat" ` //
	Name         string      `json:"name"         ` //
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
}
