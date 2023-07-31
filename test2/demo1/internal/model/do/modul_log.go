// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ModulLog is the golang structure of table modul_log for DAO operations like Where/Data.
type ModulLog struct {
	g.Meta    `orm:"table:modul_log, do:true"`
	Id        interface{} //
	ModulId   interface{} //
	Type      interface{} //
	UserUuid  interface{} //
	Url       interface{} //
	Path      interface{} //
	Size      interface{} //
	Status    interface{} //
	Time      *gtime.Time //
	StartTime *gtime.Time //
	EndTime   *gtime.Time //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
