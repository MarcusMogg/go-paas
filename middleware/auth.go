package middleware

import (
	"paas/model/entity"
	"paas/model/response"

	"github.com/gin-gonic/gin"
)

// RoleAuth 中间件判断用户是否有指定权限
// 需要先调用JWTAuth中间件
func RoleAuth(r entity.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, ok := c.Get("user")
		if !ok {
			response.FailWithMessage("未通过jwt认证", c)
			c.Abort()
			return
		}
		user := claim.(*entity.MUser)
		if user == nil || user.Role != r {
			response.FailWithMessage("权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
