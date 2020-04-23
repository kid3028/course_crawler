package data

import (
	"course_server/handler"
	"course_server/model"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/lexkong/log"
	"time"
)

var(
	// 年级科目下课程地址（专题课、系统可）
	courseDataUrl = "https://fudao.qq.com/cgi-proxy/course/discover_subject?grade=%d&subject=%d&showid=0&page=%d&size=%d"
	// 系统课地址
	pkgCourseDataUrl = "https://fudao.qq.com/cgi-proxy/course/get_course_package_info?subject_package_id=%s"
)

// 课程数据聚合
type CourseData struct {
	RetCode uint32 `json:"retcode"`
	Grade uint32 `json:"grade"`
	// 专题课
	SpecCourses SpecCourseList `json:"spe_course_list"`
	// 系统课
	CoursePkgs []CoursePkg `json:"sys_course_pkg_list"`
}

/**
  专题课列表
 */
type SpecCourseList struct {
	Page uint32 `json:"page"`
	Size uint32 `json:"size"`
	Total uint32 `json:"total"`
	// 具体课程列表
	Data []CourseInfo `json:"data"`
}

/**
	课程数据
 */
type CourseInfo struct {
	Cid uint32 `json:"cid"`
	Name string `json:"name"`
	CoverUrl string `json:"cover_url"`
	Grade uint32 `json:"grade"`
	Subject uint32 `json:"subject"`
	RecordTime int64 `json:"recordtime"`
	HasDiscount byte `json:"has_discount"`
	PreAmount uint32 `json:"pre_amount"`
	AfAmount uint32 `json:"af_amount"`
	FirstSubBgTime int64 `json:"first_sub_bgtime"`
	FirstSubEndTime int64 `json:"first_sub_endTime"`
	// 任课教师
	TeList []TeacherInfo `json:"te_list"`
	// 辅导教师
	ClassInfo struct{
		TuList []TeacherInfo
	} `json:"class_info"`
}


func (*CourseData) Tag()  {

}


/**
   根据年级科目信息抓取课程数据
 */
func FetchCourse()  {
	// 年级科目列表
	gbs, err := GradeSubjects()
	if err != nil {
		return
	}

	for _, gb := range gbs {
		c := colly.NewCollector()
		// 标记系统课处理状态，在同一个年级科目下，系统课只需存储一次，不需要分页
		initCoursePkg := true
		c.OnResponse(func(res *colly.Response) {
			var result handler.JsonResult
			if err := json.Unmarshal(res.Body, &result); err != nil {
				fmt.Println(err)
				return
			}
			var courses CourseData
			result.ParseData(&courses)
			// 保存专题课数据， 专题课为分页数据
			SaveSpeCourse(courses.SpecCourses.Data, "")
			if courses.SpecCourses.Page * courses.SpecCourses.Size < courses.SpecCourses.Total {
				res.Request.Visit(fmt.Sprintf(courseDataUrl, gb.Grade, gb.Subject, courses.SpecCourses.Page + 1,100))
			}

			// 系统课仅第一次时处理
			if initCoursePkg {
				saveCoursePkg(courses, gb.Grade, gb.Subject)
				initCoursePkg = false
			}
		})
		url := fmt.Sprintf(courseDataUrl, gb.Grade, gb.Subject, 1, 100)
		c.Request("GET", url , nil, nil, h)
	}
}

/**
  保存课程信息
 */
func SaveSpeCourse(courseInfos []CourseInfo, pkgId string) {
	if len(courseInfos) < 1 {
		return
	}

	for _, courseInfo := range courseInfos {

		course := model.SpeCourse{
			Cid:             courseInfo.Cid,
			Name:            courseInfo.Name,
			CoverUrl:        courseInfo.CoverUrl,
			Grade:           courseInfo.Grade,
			Subject:         courseInfo.Subject,
			RecordTime:      time.Unix(courseInfo.RecordTime, 0),
			HasDiscount:     courseInfo.HasDiscount,
			PreAmount:       courseInfo.PreAmount,
			AfAmount:        courseInfo.AfAmount,
			FirstSubBgTime:  time.Unix(courseInfo.FirstSubBgTime, 0),
			FirstSubEndTime: time.Unix(courseInfo.FirstSubEndTime, 0),
			CreateTime:      time.Now(),
			UpdateTime:      time.Now(),
		}

		if err := model.DB.Create(&course).Error; err != nil {
			log.Errorf(err, "save course info fail, cid: %d, name: %s, subject: %d", course.Cid, course.Name, course.Subject)
		}

		// 如果系统课程包id不为空，需要建立课程与课程包之间的关系
		if pkgId != "" {
			saveCoursePkgRelate(course.Cid, pkgId)
		}

		// 保存课程教师
		SaveTeacher(courseInfo, course)

		// 保存课程辅导
		SaveTutor(courseInfo, course)
	}

}


/**
  获取年级科目
 */
func GradeSubjects() ([]*model.GradeSubject, error) {
	gbs := make([]*model.GradeSubject, 0)

	if err := model.DB.Find(&gbs).Error; err != nil {
		log.Errorf(err, "select gradeSubject error")
		return gbs, err
	}
	return gbs, nil
}
