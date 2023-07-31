// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeat is the golang structure for table room_seat.
type RoomSeat struct {
	Id               uint        `json:"id"               ` // 逻辑id
	RoomId           int         `json:"roomId"           ` // 会议室id，room表的id
	UserUuid         uint        `json:"userUuid"         ` // 用户uuid，user表的uuid
	RoomMicrophoneId int         `json:"roomMicrophoneId" ` // 麦克风id，room_microphone表的id
	TerminalId       uint        `json:"terminalId"       ` // 终端id，terminal表的id
	LifterTerminalId int         `json:"lifterTerminalId" ` // 升降器id，terminal表的id
	ElectronicId     uint        `json:"electronicId"     ` // 电子桌牌id
	SeatNo           uint        `json:"seatNo"           ` // 座位序号
	Title            string      `json:"title"            ` // 座位描述
	Status           int         `json:"status"           ` // 状态0:停用,1启用
	Xway             int         `json:"xway"             ` // X轴
	Yway             int         `json:"yway"             ` // Y轴
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
