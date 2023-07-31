// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta               `orm:"table:user, do:true"`
	Id                   interface{} //
	Uuid                 interface{} //
	PlatformUuid         interface{} //
	Account              interface{} //
	Username             interface{} //
	Password             interface{} //
	Unit                 interface{} //
	Dept                 interface{} //
	Position             interface{} //
	Salutatory           interface{} //
	IsValid              interface{} //
	IsTmp                interface{} //
	UserType             interface{} //
	Mark                 interface{} //
	Theme                interface{} //
	Headimg              interface{} //
	Mobile               interface{} //
	Email                interface{} //
	FaceId               interface{} //
	Sort                 interface{} //
	SummitMeetingManager interface{} //
	UserLogType          interface{} //
	Deleted              interface{} //
	ChangePassword       interface{} //
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
	LastPasswdUpdateTime *gtime.Time //
	Status               interface{} //
}