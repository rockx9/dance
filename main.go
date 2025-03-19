package main

import (
	"dance/api"
	"dance/api/middleware"
	docs "dance/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func start() {
	r := gin.New()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.JWTAuthMiddleware())
	v1 := r.Group("/api/v1")
	api.Route(v1)
	r.Run("0.0.0.0:8080")
}

func main() {
	start()
}
