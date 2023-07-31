// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemFile is the golang structure for table system_file.
type SystemFile struct {
	Id            uint        `json:"id"            ` // 逻辑ID
	ParentUuid    uint        `json:"parentUuid"    ` // 父级uuid
	Uuid          uint        `json:"uuid"          ` // 文件分布式id
	PlatformUuid  uint        `json:"platformUuid"  ` // 平台uuid
	MeetingUuid   uint        `json:"meetingUuid"   ` // 会议uuid，meeting表的uuid
	UserUuid      uint        `json:"userUuid"      ` // 用户uuid，user表的uuid
	Filename      string      `json:"filename"      ` // 文件名
	Filesize      uint        `json:"filesize"      ` // 文件大小
	Filepath      string      `json:"filepath"      ` // 文件路径
	TerminalId    uint        `json:"terminalId"    ` // (t_id改为：terminal_id)终端ID：上传者终端ID 会控客户上传填 0，info_terminal（改名为terminal）表的id
	UploadTime    *gtime.Time `json:"uploadTime"    ` // 上传时间
	Sort          uint        `json:"sort"          ` // 权重
	FileUse       uint        `json:"fileUse"       ` // 文件用途:1=>议程文件，2=>议题文件，3=>纪要文件,4=>临时文件,5=>批注文件,6=>签到文件,7=>电子白板,8=>直播，9=>点播,10=>决定事项,20=>会议纪要,21=>纪要会签图片,22=>投票,23=>评分,24=>签到
	Mark          string      `json:"mark"          ` //
	IsDirectory   uint        `json:"isDirectory"   ` // 是否目录（0否1是）
	AllUser       uint        `json:"allUser"       ` // 是否选所有用户
	StartDatum    uint        `json:"startDatum"    ` // 0议题开启前显示,1议题开启后显示
	AllowDownload int         `json:"allowDownload" ` // 允许下载0不允许1允许
	IsSecret      int         `json:"isSecret"      ` // 是否保密0:不保密,1保密
	Status        int         `json:"status"        ` // 0失效
	CreatedAt     *gtime.Time `json:"createdAt"     ` // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` // 更新时间
}
