// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure for table server.
type Server struct {
	Id        uint        `json:"id"        ` // 逻辑id，服务器ID
	ServerId  int         `json:"serverId"  ` // 服务器ID
	Ip        string      `json:"ip"        ` // IP地址
	Port      uint        `json:"port"      ` // Port端口
	Status    uint        `json:"status"    ` // 状态:0离线，1=>在线
	Flag      uint        `json:"flag"      ` // 服务器标志:0=>组，1=>中心后台服务器(apache)，2=>文件服务器，3=>流媒体服务器(nginx),4=>通讯服务器(workerman)，5=>数据服务器(mysql),6=>数据服务器2(redis),7=>ps,9=>转码，10=>大屏，14=>升降器服，15=>人脸识别服务，16=>语音转写，17=>电子桌牌
	Mark      string      `json:"mark"      ` // 备注
	Pid       uint        `json:"pid"       ` // 父级id
	Operate   uint        `json:"operate"   ` // 操作状态：0=>无操作,1=>开机中，2=>关机中,3=>重启中，4=>操作成功，5=>操作失败
	Auth      uint        `json:"auth"      ` // 1=>开启、重启，2=>开启、重启、关机，3=>所有权限
	Deleted   uint        `json:"deleted"   ` // 是否删除：0否；1删除
	Name      string      `json:"name"      ` // 名称
	Content   string      `json:"content"   ` // json格式，存配置
	IsHost    int         `json:"isHost"    ` // 是否是主服务器
	UvId      int         `json:"uvId"      ` // 关联upgrade_version版本号id
	UvName    string      `json:"uvName"    ` // 守护程序上报的版本号
	UvStatus  int         `json:"uvStatus"  ` // ps版本升级状态，1=>升级中 2=>成功 3=>失败
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
