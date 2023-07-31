// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemFile is the golang structure of table system_file for DAO operations like Where/Data.
type SystemFile struct {
	g.Meta        `orm:"table:system_file, do:true"`
	Id            interface{} // 逻辑ID
	ParentUuid    interface{} // 父级uuid
	Uuid          interface{} // 文件分布式id
	PlatformUuid  interface{} // 平台uuid
	MeetingUuid   interface{} // 会议uuid，meeting表的uuid
	UserUuid      interface{} // 用户uuid，user表的uuid
	Filename      interface{} // 文件名
	Filesize      interface{} // 文件大小
	Filepath      interface{} // 文件路径
	TerminalId    interface{} // (t_id改为：terminal_id)终端ID：上传者终端ID 会控客户上传填 0，info_terminal（改名为terminal）表的id
	UploadTime    *gtime.Time // 上传时间
	Sort          interface{} // 权重
	FileUse       interface{} // 文件用途:1=>议程文件，2=>议题文件，3=>纪要文件,4=>临时文件,5=>批注文件,6=>签到文件,7=>电子白板,8=>直播，9=>点播,10=>决定事项,20=>会议纪要,21=>纪要会签图片,22=>投票,23=>评分,24=>签到
	Mark          interface{} //
	IsDirectory   interface{} // 是否目录（0否1是）
	AllUser       interface{} // 是否选所有用户
	StartDatum    interface{} // 0议题开启前显示,1议题开启后显示
	AllowDownload interface{} // 允许下载0不允许1允许
	IsSecret      interface{} // 是否保密0:不保密,1保密
	Status        interface{} // 0失效
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
