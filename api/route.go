package api

import (
	"dance/api/auth"
	"dance/api/instructor"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	auth.Route(r)
	instructor.Route(r)
}
