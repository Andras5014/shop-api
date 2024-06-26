package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_api/goods_web/models"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		adminClaims := claims.(*models.CustomClaims)
		if adminClaims.AuthorityId == 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "权限不足",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
