// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Register is the golang structure for table register.
type Register struct {
	Id        uint        `json:"id"        ` //
	Content   string      `json:"content"   ` //
	Status    uint        `json:"status"    ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
