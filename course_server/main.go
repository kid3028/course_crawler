package main

import (
	"course_server/config"
	"course_server/handler"
	"course_server/handler/course"
	"course_server/handler/task"
	"course_server/model"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var(
	cfg = pflag.StringP("config", "c", "", "crawler config file path.")
)


func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(nil)
	}
	model.InitDB()
	defer model.CloseDB()

	// 注册定时任务， 每天凌晨抓取数据
	c := cron.New()
	c.AddFunc("0 0 0 * * ?", task.Fetch)
	c.Start()
	defer c.Stop()

	router := gin.Default()
	gin.SetMode(viper.GetString("runmode"))
	router.Use(handler.Cors())

	// 路由注册
	v1 := router.Group("/v1")
	{
		v1.GET("/dayStatistic", course.StatisticSubjectByDay)
		v1.GET("/subjectCourseList", course.SubjectCourseList)
		v1.GET("/courseDetail", course.CourseDetail)
	}
	router.GET("/fetch", task.FetchHandler)

	router.Run(viper.GetString("addr"))
}
