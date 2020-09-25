package api

import (
	"fmt"
	"paas/global"
	"paas/model/entity"
	"paas/model/request"
	"paas/model/response"
	"paas/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UploadCourseFile 上传课程文件
// 需要token, cid , file
func UploadCourseFile(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	cid, err := strconv.Atoi(c.PostForm("cid"))
	if err != nil {
		response.FailValidate(c)
		return
	}
	if !service.IsCourseTeacher(uint(cid), user.ID, global.GDB) {
		response.FailWithMessage("权限不足", c)
		return
	}

	reslPath := fmt.Sprintf("source/coursefile/%d/", cid)
	fmt.Println(reslPath)
	fn, fs, err := uploadFile(reslPath, false, c)
	if err != nil {
		response.FailWithMessage("文件写入失败", c)
		return
	}
	err = service.InsertCourseFile(cid, fn+fs)
	if err != nil {
		response.FailWithMessage("数据库写入失败", c)
	} else {
		response.Ok(c)
	}
}

// GetCourseFiles 获取文件列表
func GetCourseFiles(c *gin.Context) {
	var folder request.CourseIDReq
	if err := c.BindJSON(&folder); err == nil {
		response.OkWithData(service.GetCourseFiles(folder.ID), c)
	} else {
		response.FailValidate(c)
	}
}
