// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UpgradeVersion is the golang structure for table upgrade_version.
type UpgradeVersion struct {
	Id         int         `json:"id"         ` //
	UserUuid   uint        `json:"userUuid"   ` //
	Name       string      `json:"name"       ` //
	Versions   string      `json:"versions"   ` //
	Url        string      `json:"url"        ` //
	Type       int         `json:"type"       ` //
	CurrentUse int         `json:"currentUse" ` //
	Mark       string      `json:"mark"       ` //
	Deleted    int         `json:"deleted"    ` //
	Status     int         `json:"status"     ` //
	OsType     int         `json:"osType"     ` //
	CreatedAt  *gtime.Time `json:"createdAt"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` //
}
