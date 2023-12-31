// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Meeting is the golang structure for table meeting.
type Meeting struct {
	Id               uint        `json:"id"               ` //
	PlatformUuid     uint        `json:"platformUuid"     ` //
	Uuid             uint        `json:"uuid"             ` //
	UserUuid         uint        `json:"userUuid"         ` //
	Name             string      `json:"name"             ` //
	Type             uint        `json:"type"             ` //
	Moderator        uint        `json:"moderator"        ` //
	Secretary        uint        `json:"secretary"        ` //
	StartTime        *gtime.Time `json:"startTime"        ` //
	EndTime          *gtime.Time `json:"endTime"          ` //
	RealStartTime    *gtime.Time `json:"realStartTime"    ` //
	RealEndTime      *gtime.Time `json:"realEndTime"      ` //
	MasterRoomId     int         `json:"masterRoomId"     ` //
	RoomListId       string      `json:"roomListId"       ` //
	MeetingSlogan    string      `json:"meetingSlogan"    ` //
	Nameplate        string      `json:"nameplate"        ` //
	Status           uint        `json:"status"           ` //
	IsSecrect        uint        `json:"isSecrect"        ` //
	Description      string      `json:"description"      ` //
	AllUser          int         `json:"allUser"          ` //
	ModeratorName    string      `json:"moderatorName"    ` //
	MeetingType      int         `json:"meetingType"      ` //
	IsHost           int         `json:"isHost"           ` //
	HostMeetingUuid  uint        `json:"hostMeetingUuid"  ` //
	HostPlatformUuid uint        `json:"hostPlatformUuid" ` //
	HostIp           string      `json:"hostIp"           ` //
	HostPsIp         string      `json:"hostPsIp"         ` //
	HostPsPort       int         `json:"hostPsPort"       ` //
	MenuData         string      `json:"menuData"         ` //
	MenuTab          int         `json:"menuTab"          ` //
	NameplateCustom  string      `json:"nameplateCustom"  ` //
	WelcomeCustom    string      `json:"welcomeCustom"    ` //
	CustomBgImg      string      `json:"customBgImg"      ` //
	Deleted          int         `json:"deleted"          ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
}
