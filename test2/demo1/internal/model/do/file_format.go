// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileFormat is the golang structure of table file_format for DAO operations like Where/Data.
type FileFormat struct {
	g.Meta         `orm:"table:file_format, do:true"`
	Id             interface{} //
	SystemFileUuid interface{} //
	Progress       interface{} //
	Path           interface{} //
	Status         interface{} //
	Type           interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
