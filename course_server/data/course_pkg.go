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

/**
  一个系统课程包数据
*/
type CoursePkg struct {
	PackageId         string `json:"subject_package_id"`
	Title             string `json:"title"`
	CourseBgTime      int64  `json:"course_bgtime"`
	CourseEndTime     int64  `json:"course_endtime"`
	courseSignEndTime int64  `json:"course_sign_endtime"`
	SoldCount         uint32 `json:"sold_count"`
	Season            uint32 `json:"season"`
	CourseMinPrice    uint32 `json:"course_min_price"`
	CourseMaxPrice    uint32 `json:"course_max_price"`
	DiscountPrice     uint32 `json:"discount_price"`
}

/**
  系统课程包课程列表数据
*/
type CourseOfPkg struct {
	RetCode int32         `json:"retcode"`
	Courses []*CourseInfo `json:"courses"`
}

func (*CourseOfPkg) Tag() {

}

/**
  抓取系统课课程包中内容
*/
func FetchCoursePkg() {
	pkgs := make([]model.CoursePkg, 0)
	model.DB.Find(&pkgs)
	for _, pkg := range pkgs {
		c := colly.NewCollector()
		url := fmt.Sprintf(pkgCourseDataUrl, pkg.PackageId)
		c.OnResponse(func(res *colly.Response) {
			var result handler.JsonResult
			if err := json.Unmarshal(res.Body, &result); err != nil {
				fmt.Println(err)
				return
			}
			var courses CourseOfPkg
			result.ParseData(&courses)
			// 保存课程信息，并建立课程与课程包之间的关系
			SaveSpeCourse(courses.Courses, pkg.PackageId)
		})

		c.Request("GET", url, nil, nil, h)
	}
}

/**
  保存系统课课程
*/
func saveCoursePkg(coursesData *CourseData, grade uint32, subject uint32) {
	data := coursesData.CoursePkgs
	if len(data) < 1 {
		return
	}
	for _, coursePkg := range data {
		pkg := model.CoursePkg{
			PackageId:         coursePkg.PackageId,
			Title:             coursePkg.Title,
			Grade:             grade,
			Subject:           subject,
			CourseBgTime:      time.Unix(coursePkg.CourseBgTime, 0),
			CourseEndTime:     time.Unix(coursePkg.CourseEndTime, 0),
			CourseSignEndTime: time.Unix(coursePkg.courseSignEndTime, 0),
			SoldCount:         coursePkg.SoldCount,
			Season:            coursePkg.Season,
			CourseMinPrice:    coursePkg.CourseMinPrice,
			CourseMaxPrice:    coursePkg.CourseMaxPrice,
			DiscountPrice:     coursePkg.DiscountPrice,
			CreateTime:        time.Now(),
			UpdateTime:        time.Now(),
		}

		if err := model.DB.Create(&pkg).Error; err != nil {
			log.Errorf(err, "save course pkg fail, pkgId : %s, title : %s", pkg.PackageId, pkg.Title)
		}
	}
}

/**
建立课程与系统课程包之间的关系
*/
func saveCoursePkgRelate(cid uint32, pkgId string) {
	relate := model.CoursePkgRelate{
		Cid:        cid,
		PkgId:      pkgId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err := model.DB.Create(&relate).Error; err != nil {
		log.Errorf(err, "save course pkg relation fail, cid：%s, pkgId: %s", cid, pkgId)
	}
}
