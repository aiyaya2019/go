// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ModulLog is the golang structure of table modul_log for DAO operations like Where/Data.
type ModulLog struct {
	g.Meta    `orm:"table:modul_log, do:true"`
	Id        interface{} // 主键id
	ModulId   interface{} // server的id/terminal的id
	Type      interface{} // 类型标志:0=>组，1=>中心后台服务器(apache)，2=>文件服务器，3=>流媒体服务器(nginx),4=>通讯服务器(workerman)，5=>数据服务器(mysql),6=>数据服务器2(redis),7=>ps,9=>转码，10=>大屏，14=>升降器服，15=>人脸识别服务，11=>客户端
	UserUuid  interface{} // 操作用户uuid，user表的uuid
	Url       interface{} // 下载路径
	Path      interface{} // 存放路径
	Size      interface{} // 大小
	Status    interface{} // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Time      *gtime.Time // 时间
	StartTime *gtime.Time // 开始时间
	EndTime   *gtime.Time // 结束时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
