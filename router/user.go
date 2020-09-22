package router

import (
	"paas/api"
	"paas/middleware"

	"github.com/gin-gonic/gin"
)

//InitUserRouter 初始化user路由组
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", api.Register)
		UserRouter.POST("login", api.Login)
		UserRouter.GET("myinfo", middleware.JWTAuth(), api.GetUserInfo)
		UserRouter.GET("info", api.GetUserInfoByID)
		UserRouter.POST("update", middleware.JWTAuth(), api.UpdateEmail)
		UserRouter.POST("avatar", middleware.JWTAuth(), api.UpdateAvatar)
		UserRouter.POST("password", middleware.JWTAuth(), api.UpdatePassword)
	}
}
