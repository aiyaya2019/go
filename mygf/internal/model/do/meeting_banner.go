// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingBanner is the golang structure of table meeting_banner for DAO operations like Where/Data.
type MeetingBanner struct {
	g.Meta         `orm:"table:meeting_banner, do:true"`
	Id             interface{} // 逻辑ID
	Uuid           interface{} // 标语分布式uuid
	PlatformUuid   interface{} // 平台uuid，platform表的uuid
	MeetingUuid    interface{} // 会议室uuid，meeting表的uuid
	Title          interface{} // 标题
	Background     interface{} // 背景json；status:0使用背景图/1使用背景色;color:背景色;image:背景图
	ModelFlag      interface{} // 单模式1，多模式0
	Datalayout     interface{} // 布局JSON(slogan:标语内容；slogan_above:上方标语内容；slogan_below:下方标语内容；3个内容里面的字段意思和用途一样)；status:是否显示标语内容：0不显示/1显示;left:标语框框左定位;top:标语框框顶部定位;width:标语框框宽度;height:标语框框高度;size:标语字体大小;color:字体颜色;font_id:标语字体id;font:标语字体类型;font_path:标语字体路径;text:标语内容;alignment:标语齐方式：1居中/2左对齐/3右对齐;
	Sort           interface{} // 权重
	BackgroundFlag interface{} // 应用，单:0，全:1
	BigFlag        interface{} // 非应用0，应用1
	SmallFlag      interface{} // 1推终端
	BannerImg      interface{} // 标语路径
	PushTerminal   interface{} // 推终端，默认0否
	HtmlFormatId   interface{} // html转化任务ID，html_format表的id
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
