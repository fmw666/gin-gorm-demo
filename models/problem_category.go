package models

type ProblemCategory struct {
	ID
	Timestamps
	SoftDelete
	ProblemId  int       `gorm:"column:problem_id;type:int(11);" json:"problem_id"`   // 问题 ID
	CategoryId int       `gorm:"column:category_id;type:int(11);" json:"category_id"` // 分类 ID
	Category   *Category `gorm:"foreignKey:id;references:category_id"`                // 关联分类的表
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
