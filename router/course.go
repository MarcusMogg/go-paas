package router

import (
	"paas/api"
	"paas/middleware"
	"paas/model/entity"

	"github.com/gin-gonic/gin"
)

//InitCourseRouter 初始化course路由组
func InitCourseRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("course")
	{
		UserRouter.POST("create", middleware.JWTAuth(), middleware.RoleAuth(entity.Teacher), api.CreateCourse)
		UserRouter.POST("update", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.UpdateCourse)

		UserRouter.GET("info", api.GetCourse)                                 // 根据id查找课程
		UserRouter.GET("list", api.GetCourseList)                             // 所有课程
		UserRouter.GET("mylist", middleware.JWTAuth(), api.GetCourseListByID) // 获取老师创建的所有课程或者学生加入的所有课程

		UserRouter.DELETE("del", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.DeleteCourse) // 删除课程
	}
}
