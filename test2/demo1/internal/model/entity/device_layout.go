// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceLayout is the golang structure for table device_layout.
type DeviceLayout struct {
	Id        uint        `json:"id"        ` //
	DeviceId  int         `json:"deviceId"  ` //
	RoomId    int         `json:"roomId"    ` //
	StartTime *gtime.Time `json:"startTime" ` //
	EndTime   *gtime.Time `json:"endTime"   ` //
	PushType  string      `json:"pushType"  ` //
	PushState int         `json:"pushState" ` //
	Status    int         `json:"status"    ` //
	Layout    string      `json:"layout"    ` //
	Version   string      `json:"version"   ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
