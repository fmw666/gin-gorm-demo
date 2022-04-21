package models

type Category struct {
	ID
	Timestamps
	SoftDelete
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 分类表的唯一标识
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`        // 分类名称
	ParentId int    `gorm:"column:parent_id;type:int(11);" json:"parent_id"`   // 父级 ID
}

func (table *Category) TableName() string {
	return "category"
}
