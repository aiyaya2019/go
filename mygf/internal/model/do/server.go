// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure of table server for DAO operations like Where/Data.
type Server struct {
	g.Meta    `orm:"table:server, do:true"`
	Id        interface{} // 逻辑id，服务器ID
	ServerId  interface{} // 服务器ID
	Ip        interface{} // IP地址
	Port      interface{} // Port端口
	Status    interface{} // 状态:0离线，1=>在线
	Flag      interface{} // 服务器标志:0=>组，1=>中心后台服务器(apache)，2=>文件服务器，3=>流媒体服务器(nginx),4=>通讯服务器(workerman)，5=>数据服务器(mysql),6=>数据服务器2(redis),7=>ps,9=>转码，10=>大屏，14=>升降器服，15=>人脸识别服务，16=>语音转写，17=>电子桌牌
	Mark      interface{} // 备注
	Pid       interface{} // 父级id
	Operate   interface{} // 操作状态：0=>无操作,1=>开机中，2=>关机中,3=>重启中，4=>操作成功，5=>操作失败
	Auth      interface{} // 1=>开启、重启，2=>开启、重启、关机，3=>所有权限
	Deleted   interface{} // 是否删除：0否；1删除
	Name      interface{} // 名称
	Content   interface{} // json格式，存配置
	IsHost    interface{} // 是否是主服务器
	UvId      interface{} // 关联upgrade_version版本号id
	UvName    interface{} // 守护程序上报的版本号
	UvStatus  interface{} // ps版本升级状态，1=>升级中 2=>成功 3=>失败
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
