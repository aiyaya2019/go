// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlatformGuestTask is the golang structure of table platform_guest_task for DAO operations like Where/Data.
type PlatformGuestTask struct {
	g.Meta           `orm:"table:platform_guest_task, do:true"`
	Id               interface{} // 主键
	Name             interface{} // 会议名字
	MeetingStatus    interface{} // 会议状态
	StartTime        *gtime.Time // 开始时间
	EndTime          *gtime.Time // 结束时间
	MeetingSlogan    interface{} // 会议标语
	Nameplate        interface{} // 参会人铭牌
	IsSecrect        interface{} // 会议设置(0为普通会议，1为保密会议，2开放会议)
	Description      interface{} // 会议描述
	ModeratorName    interface{} // 主席名字无关联
	CreateName       interface{} // 创建者名称
	MenuTab          interface{} // 选择的菜单数值
	AuthorityJson    interface{} // 客户端自定义菜单
	HostIp           interface{} // 主会场ip
	HostPlatformUuid interface{} // 主会场关联uuid，platform表的uuid
	HostPsIp         interface{} // 主会场主ps的ip
	HostPsPort       interface{} // 主会场主ps的port
	HostMeetingUuid  interface{} // 主会场关联会议
	Status           interface{} // 状态 1:未加入2:已加入3:加入失败4:主会场结束5:主会场删除6:被移除7:已处理8:开会中9:已退出10:已拒绝11:副平台结束
	Operate          interface{} // 操作1add 2edit 3del
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
