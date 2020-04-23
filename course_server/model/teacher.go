package model

import "time"

/**
   教师数据表实体
 */
type Teacher struct {
	Id uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Cid uint32 `gorm:"column:cid" json:"cid"`
	Name string `gorm:"column:name" json:"name"`
	CoverUrl string `gorm:"column:cover_url" json:"coverUrl"`
	Introduce string `gorm:"introduce" json:"introduce"`
	Tutor byte `gorm:"column:tutor" json:"tutor"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Teacher) TableName() string {
	return "teacher"
}