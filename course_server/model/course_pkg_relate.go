package model

import "time"

/**
   系统课课程包--课程关系表实体
 */
type CoursePkgRelate struct {
	Id uint32 `gorm:"primary key; auto_increment;column:id" json:"id"`
	Cid uint32 `gorm:"column:cid" json:"cid"`
	PkgId string `gorm:"column:pkg_id" json:"pkg_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (CoursePkgRelate) TableName() string  {
	return "course_pkg_relate"
}