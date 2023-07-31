// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MicrophoneServerTerminal is the golang structure of table microphone_server_terminal for DAO operations like Where/Data.
type MicrophoneServerTerminal struct {
	g.Meta             `orm:"table:microphone_server_terminal, do:true"`
	Id                 interface{} // 逻辑id
	MicId              interface{} // 外部传入
	Type               interface{} // 话筒类型 1.主席，0.普通
	Status             interface{} // 话筒状态 1.打开，0.关闭
	MicrophoneServerId interface{} // 连接会议主机id，microphone_server表的id
	Mark               interface{} //
	Mac                interface{} // 安卓话筒主机有带mac地址，在客户端列表关联话筒
	ScreenIp           interface{} // 安卓话筒背屏ip
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
}
