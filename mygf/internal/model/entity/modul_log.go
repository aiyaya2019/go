// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ModulLog is the golang structure for table modul_log.
type ModulLog struct {
	Id        int         `json:"id"        ` // 主键id
	ModulId   int         `json:"modulId"   ` // server的id/terminal的id
	Type      int         `json:"type"      ` // 类型标志:0=>组，1=>中心后台服务器(apache)，2=>文件服务器，3=>流媒体服务器(nginx),4=>通讯服务器(workerman)，5=>数据服务器(mysql),6=>数据服务器2(redis),7=>ps,9=>转码，10=>大屏，14=>升降器服，15=>人脸识别服务，11=>客户端
	UserUuid  uint        `json:"userUuid"  ` // 操作用户uuid，user表的uuid
	Url       string      `json:"url"       ` // 下载路径
	Path      string      `json:"path"      ` // 存放路径
	Size      int         `json:"size"      ` // 大小
	Status    int         `json:"status"    ` // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Time      *gtime.Time `json:"time"      ` // 时间
	StartTime *gtime.Time `json:"startTime" ` // 开始时间
	EndTime   *gtime.Time `json:"endTime"   ` // 结束时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
