// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Font is the golang structure of table font for DAO operations like Where/Data.
type Font struct {
	g.Meta     `orm:"table:font, do:true"`
	Id         interface{} //
	Name       interface{} //
	Path       interface{} //
	UploadTime *gtime.Time //
	Status     interface{} //
	IsSysFont  interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
