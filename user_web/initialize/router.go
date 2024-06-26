package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_api/user_web/middlewares"
	"shop_api/user_web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	// 跨域
	engine.Use(middlewares.Cors())
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	apiGroup := engine.Group("/api/v1")
	router.InitUserRouter(apiGroup)
	router.InitBaseRouter(apiGroup)

	return engine
}
