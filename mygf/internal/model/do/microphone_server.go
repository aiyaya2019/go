// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MicrophoneServer is the golang structure of table microphone_server for DAO operations like Where/Data.
type MicrophoneServer struct {
	g.Meta      `orm:"table:microphone_server, do:true"`
	Id          interface{} // 逻辑id
	RoomId      interface{} // 会议室id,room表的id
	Ip          interface{} // ip
	Port        interface{} // 端口
	MicMax      interface{} // 话筒最大打开数量 1,2,4,8
	MicType     interface{} // 话筒模式 1.FIFO 2.NORMAL 3.VOIC 4.APPLY
	Volume      interface{} // 总音量
	Alt         interface{} // 高音
	Bass        interface{} // 低音
	Sensitivity interface{} // 灵敏度
	ClosingTime interface{} // 关闭时间
	Mark        interface{} // 备注
	HostSn      interface{} // 注册码
	Status      interface{} // 1.正常  0.连接不上
	Register    interface{} // 1.已注册 0.没注册
	DeviceType  interface{} // 1. W100 2. 0200MC
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
