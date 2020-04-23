package task

import (
	"course_server/data"
	"course_server/handler"
	"course_server/model"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

var running = false

func FetchHandler(c *gin.Context) {
	go func() {
		Fetch()
	}()
	handler.SendOK(c, "抓取指令已发出")
}

func Fetch() {
	clearData()
	fetchData()
}

/**
	清空数据
 */
func clearData() {
	log.Infof("start clear old data")

	model.DB.Exec("truncate table grade_subject")
	model.DB.Exec("truncate table spe_course")
	model.DB.Exec("truncate table course_pkg")
	model.DB.Exec("truncate table course_pkg_relate")
	model.DB.Exec("truncate table teacher")

	log.Infof("complete clear old data")
}

/**
	抓取数据
 */
func fetchData() {
	log.Infof("start fetch data")

	log.Infof("start fetch gradeSubject data")
	data.FetchGradeSubject()
	log.Infof("complete fetch gradeSubject data")

	log.Infof("start fetch course data")
	data.FetchCourse()
	log.Infof("complete fetch course data")

	log.Infof("start fetch pkg course data")
	data.FetchCoursePkg()
	log.Infof("complete fetch pkg course data")

	log.Infof("complete fetch data")
}
