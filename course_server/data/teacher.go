package data

import (
	"course_server/model"
	"github.com/lexkong/log"
	"time"
)

/**
   教师
 */
type TeacherInfo struct {
	Name string `json:"name"`
	CoverUrl string `json:"cover_url"`
	Introduce string `json:"introduce"`
}

/**
  保存教师
 */
func SaveTutor(courseInfo CourseInfo, course model.SpeCourse) {
	tutors := courseInfo.ClassInfo.TuList
	if len(tutors) > 0 {
		for _, teacherInfo := range tutors {
			teacher := model.Teacher{
				Cid:        courseInfo.Cid,
				Name:       teacherInfo.Name,
				CoverUrl:   teacherInfo.CoverUrl,
				Introduce:  teacherInfo.Introduce,
				Tutor:      1,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			}
			if err := model.DB.Create(&teacher).Error; err != nil {
				log.Errorf(err, "save teacher info fail, teacher: %s, cid: %d, name: %s, subject: %d", teacher.Name, course.Cid, course.Name, course.Subject)
			}
		}
	}
}


func SaveTeacher(courseInfo CourseInfo, course model.SpeCourse) {
	if len(courseInfo.TeList) > 0 {
		for _, teacherInfo := range courseInfo.TeList {
			teacher := model.Teacher{
				Cid:        courseInfo.Cid,
				Name:       teacherInfo.Name,
				CoverUrl:   teacherInfo.CoverUrl,
				Introduce:  teacherInfo.Introduce,
				Tutor:      0,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			}
			if err := model.DB.Create(&teacher).Error; err != nil {
				log.Errorf(err, "save teacher info fail, cid: %d, name: %s, subject: %d", course.Cid, course.Name, course.Subject)
			}
		}
	}
}
