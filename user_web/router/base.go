package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/user_web/api"
	"shop_api/user_web/middlewares"
)

func InitBaseRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("base").Use(middlewares.Trace())
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)

	}
}
