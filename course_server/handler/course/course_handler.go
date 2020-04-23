package course


import (
	"course_server/handler"
	"course_server/model"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

/**
   按日统计每天各个科目新增数据
 */
func StatisticSubjectByDay(c *gin.Context)  {
	var req model.DayStatisticReq
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf(err, "请求参数解析失败")
		handler.SendErr(c, "请求参数解析失败")
		return
	}

	data := make([]model.DayStatisticSubject, 0)
	var total uint32
	model.DB.Raw("select count(distinct(`subject`)) from spe_course where record_time = ? ", req.RecordTime).Count(&total)

	// 按科目进行分组统计
	if total > 0 {
		model.DB.Raw("select record_time, `subject`, count(*) cnt from spe_course where record_time = ? group by `subject` order by cid desc", req.RecordTime).Offset((req.Current - 1) * req.Size).Limit(req.Size).Find(&data)
	}

	page := handler.Page{
		Current: req.Current,
		Size:    req.Size,
		Total:   total,
		Record: data,
	}
	handler.SendOK(c, page)

}

/**
   指定科目在指定日期新增课程列表
 */
func SubjectCourseList(c *gin.Context)  {
	var req model.SubjectCourseListReq
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf(err, "请求参数解析失败")
		handler.SendErr(c, "请求参数解析失败")
		return
	}

	data := make([]model.SpeCourse, 0)
	var total uint32
	courseDB := model.DB.Where("record_time = ? and `subject` = ?", req.RecordTime, req.Subject)
	courseDB.Model(&data).Count(&total)
	if total > 0 {
		courseDB.Offset((req.Current - 1) * req.Size).Order("cid desc").Limit(req.Size).Find(&data)
	}

	page := handler.Page{
		Current: req.Current,
		Size:    req.Size,
		Total:   total,
		Record: data,
	}
	handler.SendOK(c, page)

}

/**
  指定课程信息
 */
func CourseDetail(c *gin.Context)  {
	cid := c.Query("cid")
	if cid == "" {
		handler.SendErr(c, "未指定查询课程")
		return
	}

	var course model.SpeCourse
	model.DB.Where("cid = ? ", cid).First(&course)

	var teachers []model.Teacher
	model.DB.Where("cid = ?", cid).Find(&teachers)

	courseDetail := model.CourseDetail{
		Course:      course,
		TeacherList: teachers,
	}
	handler.SendOK(c, courseDetail)

}

