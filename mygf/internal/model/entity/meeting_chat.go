// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingChat is the golang structure for table meeting_chat.
type MeetingChat struct {
	Id          uint        `json:"id"          ` // 逻辑ID
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	UserUuid    uint        `json:"userUuid"    ` // 用户uuid，user表的uuid
	MessageType uint        `json:"messageType" ` // 消息类型：1=>聊天消息，2=> 系统消息
	Type        uint        `json:"type"        ` // 消息内容类型0=>文字，1=> 图片，2=> 语音
	Content     string      `json:"content"     ` // 消息内容
	SendTime    *gtime.Time `json:"sendTime"    ` // 发送时间
	AcceptUuid  string      `json:"acceptUuid"  ` // user表的uuid或platform表的uuid；0 代表全部uuid  多个uuid用英文半角,隔开
	Mark        string      `json:"mark"        ` // 备注
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
