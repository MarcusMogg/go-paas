package initialize

import (
	"paas/middleware"
	"paas/router"

	"github.com/gin-gonic/gin"
)

// Router 初始化路由列表
func Router() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.Cors()) // 跨域

	APIGroup := Router.Group("")
	router.InitUserRouter(APIGroup)
	router.InitCourseRouter(APIGroup)
	router.InitStudentRouter(APIGroup)
	return Router
}
