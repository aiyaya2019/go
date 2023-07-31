// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingBanner is the golang structure for table meeting_banner.
type MeetingBanner struct {
	Id             uint        `json:"id"             ` // 逻辑ID
	Uuid           uint        `json:"uuid"           ` // 标语分布式uuid
	PlatformUuid   uint        `json:"platformUuid"   ` // 平台uuid，platform表的uuid
	MeetingUuid    uint        `json:"meetingUuid"    ` // 会议室uuid，meeting表的uuid
	Title          string      `json:"title"          ` // 标题
	Background     string      `json:"background"     ` // 背景json；status:0使用背景图/1使用背景色;color:背景色;image:背景图
	ModelFlag      uint        `json:"modelFlag"      ` // 单模式1，多模式0
	Datalayout     string      `json:"datalayout"     ` // 布局JSON(slogan:标语内容；slogan_above:上方标语内容；slogan_below:下方标语内容；3个内容里面的字段意思和用途一样)；status:是否显示标语内容：0不显示/1显示;left:标语框框左定位;top:标语框框顶部定位;width:标语框框宽度;height:标语框框高度;size:标语字体大小;color:字体颜色;font_id:标语字体id;font:标语字体类型;font_path:标语字体路径;text:标语内容;alignment:标语齐方式：1居中/2左对齐/3右对齐;
	Sort           uint        `json:"sort"           ` // 权重
	BackgroundFlag uint        `json:"backgroundFlag" ` // 应用，单:0，全:1
	BigFlag        uint        `json:"bigFlag"        ` // 非应用0，应用1
	SmallFlag      int         `json:"smallFlag"      ` // 1推终端
	BannerImg      string      `json:"bannerImg"      ` // 标语路径
	PushTerminal   uint        `json:"pushTerminal"   ` // 推终端，默认0否
	HtmlFormatId   uint        `json:"htmlFormatId"   ` // html转化任务ID，html_format表的id
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
