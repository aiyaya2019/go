// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSignInfo is the golang structure of table meeting_sign_info for DAO operations like Where/Data.
type MeetingSignInfo struct {
	g.Meta         `orm:"table:meeting_sign_info, do:true"`
	Id             interface{} // 逻辑ID
	MeetingUuid    interface{} // 会议uuid，meeting表的uuid
	UserUuid       interface{} // 用户uuid，user表的uuid
	SystemFileUuid interface{} // 签到文件uuid，system_file表的uuid
	MeetingSignId  interface{} // 签到标记ID，meeting_sign表的id
	Status         interface{} // 签到状态：0未签到；1已签到
	SignTime       *gtime.Time // 签到时间
	Assist         interface{} // 是否协助签到，1:是
	DeviceType     interface{} // 设备类型：1无纸化；2电子桌牌
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
