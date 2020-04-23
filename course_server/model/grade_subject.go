package model

import "time"

/**
	年级科目数据表实体
 */
type GradeSubject struct {
	Id uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Grade uint32 `json:"grade" gorm:"column:grade"`
	Subject uint32 `json:"subject" gorm:"column:subject"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (GradeSubject) TableName() string {
	return "grade_subject"
}
