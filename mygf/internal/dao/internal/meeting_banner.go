// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingBannerDao is the data access object for table meeting_banner.
type MeetingBannerDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns MeetingBannerColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingBannerColumns defines and stores column names for table meeting_banner.
type MeetingBannerColumns struct {
	Id             string // 逻辑ID
	Uuid           string // 标语分布式uuid
	PlatformUuid   string // 平台uuid，platform表的uuid
	MeetingUuid    string // 会议室uuid，meeting表的uuid
	Title          string // 标题
	Background     string // 背景json；status:0使用背景图/1使用背景色;color:背景色;image:背景图
	ModelFlag      string // 单模式1，多模式0
	Datalayout     string // 布局JSON(slogan:标语内容；slogan_above:上方标语内容；slogan_below:下方标语内容；3个内容里面的字段意思和用途一样)；status:是否显示标语内容：0不显示/1显示;left:标语框框左定位;top:标语框框顶部定位;width:标语框框宽度;height:标语框框高度;size:标语字体大小;color:字体颜色;font_id:标语字体id;font:标语字体类型;font_path:标语字体路径;text:标语内容;alignment:标语齐方式：1居中/2左对齐/3右对齐;
	Sort           string // 权重
	BackgroundFlag string // 应用，单:0，全:1
	BigFlag        string // 非应用0，应用1
	SmallFlag      string // 1推终端
	BannerImg      string // 标语路径
	PushTerminal   string // 推终端，默认0否
	HtmlFormatId   string // html转化任务ID，html_format表的id
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// meetingBannerColumns holds the columns for table meeting_banner.
var meetingBannerColumns = MeetingBannerColumns{
	Id:             "id",
	Uuid:           "uuid",
	PlatformUuid:   "platform_uuid",
	MeetingUuid:    "meeting_uuid",
	Title:          "title",
	Background:     "background",
	ModelFlag:      "model_flag",
	Datalayout:     "datalayout",
	Sort:           "sort",
	BackgroundFlag: "background_flag",
	BigFlag:        "big_flag",
	SmallFlag:      "small_flag",
	BannerImg:      "banner_img",
	PushTerminal:   "push_terminal",
	HtmlFormatId:   "html_format_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewMeetingBannerDao creates and returns a new DAO object for table data access.
func NewMeetingBannerDao() *MeetingBannerDao {
	return &MeetingBannerDao{
		group:   "default",
		table:   "meeting_banner",
		columns: meetingBannerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingBannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingBannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingBannerDao) Columns() MeetingBannerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingBannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingBannerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingBannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
