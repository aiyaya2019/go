// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Terminal is the golang structure for table terminal.
type Terminal struct {
	Id                      uint        `json:"id"                      ` // 逻辑id
	TId                     uint        `json:"tId"                     ` // 终端id
	Name                    string      `json:"name"                    ` // 终端名称
	Ip                      string      `json:"ip"                      ` // 终端IP
	Mac                     string      `json:"mac"                     ` // 终端 mac
	Status                  uint        `json:"status"                  ` // 状态:0=>离线，1=>在线
	RoomId                  uint        `json:"roomId"                  ` // 会议室ID，room表的id
	Mark                    string      `json:"mark"                    ` // 备注
	Type                    uint        `json:"type"                    ` // 1.pc客户端2.移动客户端4.解码器大屏8.升降器终端,9三代升降器,10二代升降器,关联二代8终端
	Config                  string      `json:"config"                  ` // 单个终端配置
	ConfigStatus            int         `json:"configStatus"            ` // 终端下发配置，0=>默认1=>成功2=>失败
	AssociatedVersionId     int         `json:"associatedVersionId"     ` // 关联版本id
	AssociatedVersionName   string      `json:"associatedVersionName"   ` // 版本名称
	AssociatedVersionStatus int         `json:"associatedVersionStatus" ` // 升级状态，1=>升级中 2=>成功 3=>失败
	Operate                 int         `json:"operate"                 ` // 操作状态：0=>无操作,1=>开机中，2=>关机中,3=>重启中，4=>操作成功，5=>操作失败，6=>操作异常
	OsType                  int         `json:"osType"                  ` // 1：windows；2：android; 3: ios;4: kylin； 5：centos;
	GuardConfig             string      `json:"guardConfig"             ` // 守护程序配置
	GuardUvId               int         `json:"guardUvId"               ` // 守护程序版本id
	GuardUvName             string      `json:"guardUvName"             ` // 守护程序版本名称
	GuardUvStatus           int         `json:"guardUvStatus"           ` // 守护程序版本状态
	DeviceSn                string      `json:"deviceSn"                ` // 序列号
	DeviceType              int         `json:"deviceType"              ` // 型号
	IpType                  int         `json:"ipType"                  ` // IP类型0动态，1静态
	Mask                    string      `json:"mask"                    ` // 子网掩码
	Gateway                 string      `json:"gateway"                 ` // 默认网关
	HostIp                  string      `json:"hostIp"                  ` // 主机ip
	ClientIp                string      `json:"clientIp"                ` // 客户端ip
	OptCounts               int         `json:"optCounts"               ` // 上升下降次数
	Concurrence             string      `json:"concurrence"             ` // 并发用户名
	LifterType              int         `json:"lifterType"              ` // 终端升降器类型2代3代
	SecondTid               int         `json:"secondTid"               ` // type10 二代升降器关联type8的id
	CreatedAt               *gtime.Time `json:"createdAt"               ` // 创建时间
	UpdatedAt               *gtime.Time `json:"updatedAt"               ` // 更新时间
}
