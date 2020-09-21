package router

import (
	"paas/api"
	"paas/middleware"

	"github.com/gin-gonic/gin"
)

//InitStudentRouter 学生管理相关
func InitStudentRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("student")
	{
		UserRouter.POST("add", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.AddStudents)
		UserRouter.GET("list", api.GetStudents)
		UserRouter.DELETE("del", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.DeleteStudent)
	}
}
