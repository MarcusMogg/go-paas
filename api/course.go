package api

import (
	"paas/model/entity"
	"paas/model/request"
	"paas/model/response"
	"paas/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCourse 创建课程
func CreateCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var course request.CourseReq
	if err := c.BindJSON(&course); err == nil {
		courseData := entity.Course{
			TeacherID:   user.ID,
			Instruction: course.Instruction,
			Name:        course.Name,
		}
		if err = service.InsertCourse(&courseData, user.ID); err == nil {
			response.Ok(c)
		} else {
			response.FailWithMessage("课程创建失败", c)
		}
	} else {
		response.FailValidate(c)
	}

}

// UpdateCourse 修改课程信息
func UpdateCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var cs request.CourseUReq
	if err := c.BindJSON(&cs); err == nil {
		course := service.GetCourseByID(cs.ID)
		course.Instruction = cs.Instruction
		course.Name = cs.Name
		if err = service.UpdateCourse(course, user); err == nil {
			response.Ok(c)
		} else {
			response.FailWithMessage(err.Error(), c)
		}
	} else {
		response.FailValidate(c)
	}

}

// GetCourse 读取课程信息
func GetCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	course := service.GetCourseByID(uint(id))
	response.OkWithData(course, c)
}

// GetCourseListByID 读取教师创建的课程列表
func GetCourseListByID(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	courses := service.GetMyCourses(user.ID)
	response.OkWithData(courses, c)
}

// GetCourseList 获取所有课程
func GetCourseList(c *gin.Context) {
	pagenum, err1 := strconv.Atoi(c.DefaultQuery("pagenum", "1"))
	pagesize, err2 := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	keyword := c.Query("key")
	if err1 != nil || err2 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	courses, tot := service.GetCourses(pagenum, pagesize, keyword)
	response.OkWithData(gin.H{
		"total": tot,
		"cs":    courses,
	}, c)
}

// DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		service.DropCourse(id.ID, user.ID)
	} else {
		response.FailValidate(c)
	}
}
