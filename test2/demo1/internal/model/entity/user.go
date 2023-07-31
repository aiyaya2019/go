// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                   uint        `json:"id"                   ` //
	Uuid                 uint        `json:"uuid"                 ` //
	PlatformUuid         uint        `json:"platformUuid"         ` //
	Account              string      `json:"account"              ` //
	Username             string      `json:"username"             ` //
	Password             string      `json:"password"             ` //
	Unit                 string      `json:"unit"                 ` //
	Dept                 string      `json:"dept"                 ` //
	Position             string      `json:"position"             ` //
	Salutatory           string      `json:"salutatory"           ` //
	IsValid              uint        `json:"isValid"              ` //
	IsTmp                int         `json:"isTmp"                ` //
	UserType             uint        `json:"userType"             ` //
	Mark                 string      `json:"mark"                 ` //
	Theme                uint        `json:"theme"                ` //
	Headimg              string      `json:"headimg"              ` //
	Mobile               string      `json:"mobile"               ` //
	Email                string      `json:"email"                ` //
	FaceId               int         `json:"faceId"               ` //
	Sort                 int         `json:"sort"                 ` //
	SummitMeetingManager int         `json:"summitMeetingManager" ` //
	UserLogType          int         `json:"userLogType"          ` //
	Deleted              int         `json:"deleted"              ` //
	ChangePassword       int         `json:"changePassword"       ` //
	CreatedAt            *gtime.Time `json:"createdAt"            ` //
	UpdatedAt            *gtime.Time `json:"updatedAt"            ` //
	LastPasswdUpdateTime *gtime.Time `json:"lastPasswdUpdateTime" ` //
	Status               int         `json:"status"               ` //
}