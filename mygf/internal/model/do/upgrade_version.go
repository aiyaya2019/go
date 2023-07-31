// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpgradeVersion is the golang structure of table upgrade_version for DAO operations like Where/Data.
type UpgradeVersion struct {
	g.Meta     `orm:"table:upgrade_version, do:true"`
	Id         interface{} //
	UserUuid   interface{} // 上传用户uuid，user表的uuid
	Name       interface{} // 文件名
	Versions   interface{} // 版本
	Url        interface{} // 文件地址
	Type       interface{} // 4=>cms,5=>mysql,7=>ps,10=>大屏,11=>client,12=>守护
	CurrentUse interface{} // cms、client、mysql使用版本
	Mark       interface{} // 备注
	Deleted    interface{} // 1=>删除
	Status     interface{} // 1=>升级中 2=>成功 3=>失败
	OsType     interface{} // 1：windows；2：android; 3: ios;4: kylin； 5：centos;
	CreatedAt  *gtime.Time // 创建时间/上传时间
	UpdatedAt  *gtime.Time // 更新时间
}
