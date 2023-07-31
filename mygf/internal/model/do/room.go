// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Room is the golang structure of table room for DAO operations like Where/Data.
type Room struct {
	g.Meta         `orm:"table:room, do:true"`
	Id             interface{} // 逻辑id,会议室ID
	Name           interface{} // 会议室名称
	ServerId       interface{} // 服务器ID，server表的id
	Roomlayout     interface{} // 会议室布局json(单个)
	Status         interface{} // 状态:0=>未使用，1=>使用中
	Mark           interface{} // 备注
	Roombackground interface{} // 会议室布局背景
	BgWidth        interface{} // 背景图宽
	BgHeight       interface{} // 背景图高
	TerminalConfig interface{} // 终端配置json格式
	Transfer       interface{} // 0不启用语音转写，1启用
	Electronic     interface{} // 0不启用电子桌牌，1启用
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	TransferRoomId interface{} //
}
