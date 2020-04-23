package model

import (
"time"
)

/**
  课程数据表实体
 */
type SpeCourse struct {
	Cid uint32 `gorm:"primary_key column:cid" json:"cid"`
	Name string `gorm:"column:name" json:"name"`
	CoverUrl string `gorm:"column:cover_url" json:"coverUrl"`
	Grade uint32 `gorm:"column:grade" json:"grade"`
	Subject uint32 `gorm:"subject" json:"subject"`
	RecordTime time.Time `gorm:"record_time" json:"recordTime"`
	HasDiscount byte `gorm:"has_discount" json:"hasDiscount"`
	PreAmount uint32 `gorm:"pre_amount" json:"preAmount"`
	AfAmount uint32 `gorm:"af_amount" json:"afAmount"`
	FirstSubBgTime time.Time `gorm:"first_sub_bg_time" json:"firstSubBgTime"`
	FirstSubEndTime time.Time `gorm:"first_sub_end_time" json:"firstSubEndTime"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}


func (SpeCourse) TableName() string  {
	return "spe_course"
}

/**
   日新增课程科目 响应模型
 */
type DayStatisticSubject struct {
	RecordTime string `gorm:"column:record_time" json:"recordTime"`
	Subject uint32 `gorm:"subject" json:"subject"`
	Cnt uint32 `gorm:"cnt" json:"cnt"`
}

/**
    日新增课程科目 请求模型
 */
type DayStatisticReq struct {
	RecordTime string `form:"recordTime" binding:"required"`
	Current uint32 `form:"current" binding:"required"`
	Size uint32 `form:"size" binding:"required"`
}

/**
  	科目课程列表 请求模型
 */
type SubjectCourseListReq struct {
	RecordTime string `form:"recordTime" binding:"required"`
	Subject uint32 `form:"subject" binding:"required"`
	Current uint32 `form:"current" binding:"required"`
	Size uint32 `form:"size" binding:"required"`
}

/**
   课程详情 响应模型
 */
type CourseDetail struct {
	Course SpeCourse `json:"course"`
	TeacherList []Teacher `json:"teacherList"`
}





