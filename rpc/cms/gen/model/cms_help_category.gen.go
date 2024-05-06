// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCmsHelpCategory = "cms_help_category"

// CmsHelpCategory 帮助分类表
type CmsHelpCategory struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	Icon       string `gorm:"column:icon;not null;comment:分类图标" json:"icon"`             // 分类图标
	HelpCount  int32  `gorm:"column:help_count;not null;comment:专题数量" json:"help_count"` // 专题数量
	ShowStatus int32  `gorm:"column:show_status;not null" json:"show_status"`
	Sort       int32  `gorm:"column:sort;not null" json:"sort"`
}

// TableName CmsHelpCategory's table name
func (*CmsHelpCategory) TableName() string {
	return TableNameCmsHelpCategory
}
