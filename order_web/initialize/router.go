package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_api/order_web/middlewares"
	"shop_api/order_web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 跨域
	engine.Use(middlewares.Cors())
	//添加链路追踪
	ApiGroup := engine.Group("/o/v1")
	router.InitOrderRouter(ApiGroup)
	router.InitShopCartRouter(ApiGroup)

	return engine
}
