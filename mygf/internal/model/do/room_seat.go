// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeat is the golang structure of table room_seat for DAO operations like Where/Data.
type RoomSeat struct {
	g.Meta           `orm:"table:room_seat, do:true"`
	Id               interface{} // 逻辑id
	RoomId           interface{} // 会议室id，room表的id
	UserUuid         interface{} // 用户uuid，user表的uuid
	RoomMicrophoneId interface{} // 麦克风id，room_microphone表的id
	TerminalId       interface{} // 终端id，terminal表的id
	LifterTerminalId interface{} // 升降器id，terminal表的id
	ElectronicId     interface{} // 电子桌牌id
	SeatNo           interface{} // 座位序号
	Title            interface{} // 座位描述
	Status           interface{} // 状态0:停用,1启用
	Xway             interface{} // X轴
	Yway             interface{} // Y轴
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
