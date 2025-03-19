package auth

import "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {
	api := r.Group("/auth")
	api.POST("/login", LoginHandle)
}
