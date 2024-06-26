package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_api/goods_web/middlewares"
	"shop_api/goods_web/router"
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
	ApiGroup := engine.Group("/g/v1")
	router.InitGoodsRouter(ApiGroup)
	//router.InitCategoryRouter(ApiGroup)
	//router.InitBannerRouter(ApiGroup)
	//router.InitBrandRouter(ApiGroup)

	return engine
}
