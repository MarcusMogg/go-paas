package router

import (
	"paas/api"
	"paas/middleware"

	"github.com/gin-gonic/gin"
)

//InitDockerRouter 初始化course路由组
func InitDockerRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("docker")
	{
		UserRouter.POST("pull", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.DockerPull)
		UserRouter.POST("images", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.GetDockerImages)
		UserRouter.POST("del", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.DelDockerImage)

		UserRouter.POST("create", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.CreateContainer)
		UserRouter.POST("alloc", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.AllocContainer)
		UserRouter.POST("containerlist", middleware.JWTAuth(), middleware.CourseAuth(), api.GetContainerList)
		UserRouter.POST("container", api.GetContainer)
		UserRouter.GET("exec", api.Terminal)
		UserRouter.POST("delall", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.DropAllContainer)
		UserRouter.POST("delc", middleware.JWTAuth(), middleware.CourseTeacherAuth(), api.DropContainer)
	}
}
