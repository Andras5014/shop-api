package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/user_web/api"
)

func InitBaseRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)

	}
}
