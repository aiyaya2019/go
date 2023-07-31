// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingChat is the golang structure of table meeting_chat for DAO operations like Where/Data.
type MeetingChat struct {
	g.Meta      `orm:"table:meeting_chat, do:true"`
	Id          interface{} // 逻辑ID
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	UserUuid    interface{} // 用户uuid，user表的uuid
	MessageType interface{} // 消息类型：1=>聊天消息，2=> 系统消息
	Type        interface{} // 消息内容类型0=>文字，1=> 图片，2=> 语音
	Content     interface{} // 消息内容
	SendTime    *gtime.Time // 发送时间
	AcceptUuid  interface{} // user表的uuid或platform表的uuid；0 代表全部uuid  多个uuid用英文半角,隔开
	Mark        interface{} // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
