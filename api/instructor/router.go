package instructor

import "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {
	api := r.Group("/instructors")
	api.GET("", GetInstructorListHandle)
	api.POST("", CreateInstructorHandle)
	api.GET("/:id", GetInstructorHandle)
	api.PUT("/:id", UpdateInstructorHandle)
	api.DELETE("/:id", DeleteInstructorHandle)
}
