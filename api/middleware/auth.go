package middleware

import (
	"dance/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 基于JWT认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.Request.URL.Path, "/auth/login") && c.Request.Method == http.MethodPost {
			c.Next()
			return
		}
		authHeader := c.Request.Header.Get("Authorization")
		user, err := service.NewAuth().ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
				"data": nil,
			})
			c.Abort()
			return
		}
		now := time.Now()
		if now.Before(user.NotBefore.Time) || now.After(user.ExpiresAt.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "Expired token.",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Set("username", user.Username)
		c.Next()
	}

}
