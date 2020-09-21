package api

import (
	"paas/model/entity"
	"paas/model/request"
	"paas/model/response"
	"paas/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddStudents 批量添加学生
func AddStudents(c *gin.Context) {
	var as request.AddStudentsReq
	if err := c.BindJSON(&as); err == nil {
		var errs []string
		for _, i := range as.UserNames {
			if _, err := service.InsertStudent(as.ID, i, 1); err != nil {
				errs = append(errs, i)
			}
		}
		if len(errs) != 0 {
			response.FailDetailed(response.ERROR, errs, "用户名错误", c)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}

// GetStudents 获取学生列表
func GetStudents(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	var user []entity.MUser = service.GetStudents(uint(id))
	response.OkWithData(user, c)

}

// DeleteStudent 删除学生
func DeleteStudent(c *gin.Context) {
	var a request.DelStudentReq
	if err := c.BindJSON(&a); err == nil {
		err := service.DeleteStudent(a.GetByID.ID, a.CourseIDReq.ID)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.Ok(c)
		}
	} else {
		response.FailValidate(c)
	}
}
