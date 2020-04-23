package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  爬虫数据统一影响
 */
type JsonResult struct {
	RetCode int32 `json:"retcode"`
	Result json.RawMessage `json:"result"`
}

type IBean interface{ Tag() }

/**
   转换响应data到指定实体上
 */
func (data *JsonResult) ParseData(bean IBean)  {
	if err := json.Unmarshal([]byte(data.Result), bean); err != nil {
		fmt.Println(err)
	}
}

/**
   web统一响应
 */
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/**
   分页数据模型
 */
type Page struct {
	Current uint32 `json:"current"`
	Size uint32 `json:"size"`
	Total uint32 `json:"total"`
	Record interface{} `json:"record"`
}

/**
   成功响应
 */
func SendOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Data: data,
	})
}

/**
	失败响应
 */
func SendErr(c *gin.Context, msg string) {

	c.JSON(http.StatusOK, Response{
		Code:    1,
		Message: msg,
	})
}