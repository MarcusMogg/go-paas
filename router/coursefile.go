package router

import (
	"paas/api"
	"paas/middleware"

	"github.com/gin-gonic/gin"
)

//InitCourseFileRouter 初始化course路由组
func InitCourseFileRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("coursefile")
	{
		UserRouter.POST("createfile", middleware.JWTAuth(), api.UploadCourseFile)
		UserRouter.POST("files", middleware.JWTAuth(), middleware.CourseAuth(), api.GetCourseFiles)
	}
}
