// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UpgradeVersion is the golang structure for table upgrade_version.
type UpgradeVersion struct {
	Id         int         `json:"id"         ` //
	UserUuid   uint        `json:"userUuid"   ` // 上传用户uuid，user表的uuid
	Name       string      `json:"name"       ` // 文件名
	Versions   string      `json:"versions"   ` // 版本
	Url        string      `json:"url"        ` // 文件地址
	Type       int         `json:"type"       ` // 4=>cms,5=>mysql,7=>ps,10=>大屏,11=>client,12=>守护
	CurrentUse int         `json:"currentUse" ` // cms、client、mysql使用版本
	Mark       string      `json:"mark"       ` // 备注
	Deleted    int         `json:"deleted"    ` // 1=>删除
	Status     int         `json:"status"     ` // 1=>升级中 2=>成功 3=>失败
	OsType     int         `json:"osType"     ` // 1：windows；2：android; 3: ios;4: kylin； 5：centos;
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间/上传时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 更新时间
}
