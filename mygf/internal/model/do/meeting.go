// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Meeting is the golang structure of table meeting for DAO operations like Where/Data.
type Meeting struct {
	g.Meta           `orm:"table:meeting, do:true"`
	Id               interface{} // 会议ID
	PlatformUuid     interface{} // 平台uuid，platform表的uuid
	Uuid             interface{} // 会议分布式uuid
	UserUuid         interface{} // 创建者用户uuid，user表的uuid
	Name             interface{} // 会议名称
	Type             interface{} // 类型:0=>不要登录，1=>要登录
	Moderator        interface{} // 会议主持人，user表的uuid
	Secretary        interface{} // 会议秘书
	StartTime        *gtime.Time // 开始时间
	EndTime          *gtime.Time // 结束时间
	RealStartTime    *gtime.Time // 真正开始时间
	RealEndTime      *gtime.Time // 真正结束时间
	MasterRoomId     interface{} // 主会议室id,room表的id
	RoomListId       interface{} // 会议室ID;多个会议室ID用’,’号隔开
	MeetingSlogan    interface{} // 会议标语
	Nameplate        interface{} // 参会人铭牌
	Status           interface{} // 状态:0未开始，1开会中，2结束,3模板会议,4暂停
	IsSecrect        interface{} // 会议设置(0为普通会议，1为保密会议，2开放会议)
	Description      interface{} // 会议描述
	AllUser          interface{} // 0手选,1全选
	ModeratorName    interface{} // 主席名字无关联
	MeetingType      interface{} // 会议类型,1常规2本地3异地
	IsHost           interface{} // 1.主会场2.客会场
	HostMeetingUuid  interface{} // 主会场关联uuid，meeting表的uuid
	HostPlatformUuid interface{} // 主会场平台uuid，platform表的uuid
	HostIp           interface{} // 主会场ip
	HostPsIp         interface{} // 主会场主ps的ip
	HostPsPort       interface{} // 主会场主ps的port
	MenuData         interface{} // 会议菜单信息
	MenuTab          interface{} // 菜单模式：1卡片模式；2简洁模式;3导航模式功能菜单；4经典模式功能菜单
	NameplateCustom  interface{} // 铭牌的自定义文本，限制50个汉字
	WelcomeCustom    interface{} // 欢迎界面的自定义文本，限制50个汉字
	CustomBgImg      interface{} // 自定义背景图
	Deleted          interface{} // 删除：0未删除；1删除到回收站；2删除
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
