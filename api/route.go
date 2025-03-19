package api

import (
	"dance/api/instructor"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	instructor.Route(r)
}
