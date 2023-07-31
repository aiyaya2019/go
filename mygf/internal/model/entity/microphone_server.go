// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MicrophoneServer is the golang structure for table microphone_server.
type MicrophoneServer struct {
	Id          uint        `json:"id"          ` // 逻辑id
	RoomId      uint        `json:"roomId"      ` // 会议室id,room表的id
	Ip          string      `json:"ip"          ` // ip
	Port        uint        `json:"port"        ` // 端口
	MicMax      uint        `json:"micMax"      ` // 话筒最大打开数量 1,2,4,8
	MicType     int         `json:"micType"     ` // 话筒模式 1.FIFO 2.NORMAL 3.VOIC 4.APPLY
	Volume      int         `json:"volume"      ` // 总音量
	Alt         int         `json:"alt"         ` // 高音
	Bass        int         `json:"bass"        ` // 低音
	Sensitivity int         `json:"sensitivity" ` // 灵敏度
	ClosingTime int         `json:"closingTime" ` // 关闭时间
	Mark        string      `json:"mark"        ` // 备注
	HostSn      string      `json:"hostSn"      ` // 注册码
	Status      int         `json:"status"      ` // 1.正常  0.连接不上
	Register    int         `json:"register"    ` // 1.已注册 0.没注册
	DeviceType  int         `json:"deviceType"  ` // 1. W100 2. 0200MC
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
