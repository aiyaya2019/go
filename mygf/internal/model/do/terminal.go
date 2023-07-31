// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Terminal is the golang structure of table terminal for DAO operations like Where/Data.
type Terminal struct {
	g.Meta                  `orm:"table:terminal, do:true"`
	Id                      interface{} // 逻辑id
	TId                     interface{} // 终端id
	Name                    interface{} // 终端名称
	Ip                      interface{} // 终端IP
	Mac                     interface{} // 终端 mac
	Status                  interface{} // 状态:0=>离线，1=>在线
	RoomId                  interface{} // 会议室ID，room表的id
	Mark                    interface{} // 备注
	Type                    interface{} // 1.pc客户端2.移动客户端4.解码器大屏8.升降器终端,9三代升降器,10二代升降器,关联二代8终端
	Config                  interface{} // 单个终端配置
	ConfigStatus            interface{} // 终端下发配置，0=>默认1=>成功2=>失败
	AssociatedVersionId     interface{} // 关联版本id
	AssociatedVersionName   interface{} // 版本名称
	AssociatedVersionStatus interface{} // 升级状态，1=>升级中 2=>成功 3=>失败
	Operate                 interface{} // 操作状态：0=>无操作,1=>开机中，2=>关机中,3=>重启中，4=>操作成功，5=>操作失败，6=>操作异常
	OsType                  interface{} // 1：windows；2：android; 3: ios;4: kylin； 5：centos;
	GuardConfig             interface{} // 守护程序配置
	GuardUvId               interface{} // 守护程序版本id
	GuardUvName             interface{} // 守护程序版本名称
	GuardUvStatus           interface{} // 守护程序版本状态
	DeviceSn                interface{} // 序列号
	DeviceType              interface{} // 型号
	IpType                  interface{} // IP类型0动态，1静态
	Mask                    interface{} // 子网掩码
	Gateway                 interface{} // 默认网关
	HostIp                  interface{} // 主机ip
	ClientIp                interface{} // 客户端ip
	OptCounts               interface{} // 上升下降次数
	Concurrence             interface{} // 并发用户名
	LifterType              interface{} // 终端升降器类型2代3代
	SecondTid               interface{} // type10 二代升降器关联type8的id
	CreatedAt               *gtime.Time // 创建时间
	UpdatedAt               *gtime.Time // 更新时间
}
