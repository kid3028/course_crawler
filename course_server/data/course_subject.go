package data

import (
	"course_server/handler"
	"course_server/model"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/lexkong/log"
	"net/http"
	"sync"
	"time"
)

var(

	// 设置头部
	h = http.Header{
		"Referer": []string{"https://fudao.qq.com/"},
	}
	// grade  --> subject url
	gradeSubjectUrl = "https://fudao.qq.com/cgi-proxy/course/grade_subject"

)


type GradeSubjectList struct {
	GradeSubjects []GradeSubjectData `json:"grade_subjects"`
}

/**
  科目年级
 */
type GradeSubjectData struct {
	Grade   uint32   `json:"grade"`
	Subject []uint32 `json:"subject"`
}

func (*GradeSubjectList) Tag()  {

}

/**
   启动科目年级抓取
 */
func FetchGradeSubject() {
	c := colly.NewCollector()

	c.OnResponse(parseGradeSubjects)
	c.Request("GET", gradeSubjectUrl, nil, nil, h)
}

/**
	处理科目年级抓取数据
 */
func parseGradeSubjects(res *colly.Response)  {
	// 反序列抓取对象到公共响应
	var result handler.JsonResult
	if err := json.Unmarshal(res.Body, &result); err != nil {
		fmt.Println(err)
		return
	}

	if result.RetCode != 0 {
		fmt.Println("拉取年级科目失败")
		return
	}

	// 反序列化响应数据到科目年级对象
	var gradeSubjects GradeSubjectList
	result.ParseData(&gradeSubjects)

	// 持久到数据库
	wg := sync.WaitGroup{}
	for _, gradeSubject := range gradeSubjects.GradeSubjects {
		wg.Add(1)
		go SaveGradeSubject(gradeSubject, &wg)
	}
	wg.Wait()
}

/**
	保存年级科目关系
 */
func SaveGradeSubject(data GradeSubjectData, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, subject := range data.Subject {
		gb := model.GradeSubject{
			Grade:    data.Grade,
			Subject:  subject,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}

		err := model.DB.Exec("insert into grade_subject(grade, subject, create_time, update_time) select ?, ?, ?, ? from dual where not exists (select 1 from grade_subject where grade = ? and subject = ?)",
			gb.Grade, gb.Subject, gb.CreateTime, gb.UpdateTime, gb.Grade, gb.Subject).Error
		if err != nil {
			log.Errorf(err, "save gradeSubject fail, grade:%d, subject : %d", gb.Grade, gb.Subject)
		}
	}
}
