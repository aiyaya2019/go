// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingService is the golang structure of table meeting_service for DAO operations like Where/Data.
type MeetingService struct {
	g.Meta              `orm:"table:meeting_service, do:true"`
	Id                  interface{} // 逻辑ID
	MeetingUuid         interface{} // 会议uuid，meeting表的uuid
	UserUuid            interface{} // 用户uuid，user表的uuid
	ServerId            interface{} // 服务ID，server表的id
	Content             interface{} // 服务内容
	Time                *gtime.Time // 时间
	IsConfirm           interface{} // 是否确认：0=>未确认，1=> 已确认
	Mark                interface{} // 备注
	Status              interface{} // 0:未处理，1:稍后处理，2:已处理，3:已过期,4:已取消
	LaterHandleUserUuid interface{} // 稍后处理者uuid，user表的uuid
	HandleUserUuid      interface{} // 处理者uuid，user表的uuid
	LaterHandleTime     *gtime.Time // 稍后处理时间
	HandleTime          *gtime.Time // 处理时间
	DeviceType          interface{} // 设备类型：1无纸化；2电子桌牌
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
}
