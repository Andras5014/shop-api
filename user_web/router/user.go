package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shop_api/user_web/api"
	"shop_api/user_web/middlewares"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	zap.S().Info("初始化用户模块的路由信息")

	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
		UserRouter.POST("register", api.Register)
	}
}
