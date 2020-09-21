package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"paas/global"
	"paas/model/entity"
	"paas/model/request"
	"paas/model/response"
	"paas/service"

	"github.com/gin-gonic/gin"
)

var errNoCid = errors.New("未携带cid")

// CourseAuth 中间件判断用户是否属于课程
// 需要先调用JWTAuth中间件
// 传入参数必须包含CourseIDReq
func CourseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("user")
		if !ok {
			response.FailWithMessage("未通过jwt认证", c)
			c.Abort()
			return
		}
		user := claim.(*entity.MUser)
		if cid, err := getCid(c); err == nil {
			if service.InCourse(cid, user.ID, global.GDB) {
				c.Next()
				return
			}
			response.FailWithMessage("不属于课程", c)
			c.Abort()
		}
	}
}

// CourseTeacherAuth 中间件判断用户是否是创建课程的老师
// 需要先调用JWTAuth中间件
// 传入参数必须包含CourseIDReq
func CourseTeacherAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("user")
		if !ok {
			response.FailWithMessage("未通过jwt认证", c)
			c.Abort()
			return
		}
		user := claim.(*entity.MUser)
		if cid, err := getCid(c); err == nil {
			if service.IsCourseTeacher(cid, user.ID, global.GDB) {
				c.Next()
				return
			}
			response.FailWithMessage("不是课程教师", c)
			c.Abort()
		}
	}
}

func getCid(c *gin.Context) (uint, error) {
	if data, err := c.GetRawData(); err == nil {
		var cid request.CourseIDReq
		if err := json.Unmarshal(data, &cid); err == nil {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			return cid.ID, nil
		}
	}
	response.FailWithMessage(errNoCid.Error(), c)
	c.Abort()
	return 0, errNoCid
}
