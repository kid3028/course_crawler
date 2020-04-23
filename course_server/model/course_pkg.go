package model

import "time"

/**
   系统课课程包数据表实体
 */
type CoursePkg struct {
	PackageId string `gorm:"primary key;column:package_id" json:"packageId"`
	Title string `gorm:"column:title" json:"title"`
	Grade uint32 `gorm:"column:grade" json:"grade"`
	Subject uint32 `gorm:"column:subject" json:"subject"`
	CourseBgTime time.Time `gorm:"column:course_bg_time" json:"courseBgTime" `
	CourseEndTime time.Time `gorm:"column:course_end_time" json:"courseEndTime"`
	CourseSignEndTime time.Time `gorm:"column:course_sign_end_time" json:"courseSignEndTime"`
	SoldCount uint32 `gorm:"column:sold_count" json:"soldCount"`
	Season uint32 `gorm:"column:season" json:"season"`
	CourseMinPrice uint32 `gorm:"column:course_min_price" json:"courseMinPrice"`
	CourseMaxPrice uint32 `gorm:"column:course_max_price" json:"courseMaxPrice"`
	DiscountPrice uint32 `gorm:"column:discount_price" json:"discountPrice"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (CoursePkg) TableName() string  {
	return "course_pkg"

}