package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	GC        *gin.Context
	RequestId string
	User      string
	Company   string
}

func (c *Context) SetGinContext(gc *gin.Context) {
	c.GC = gc
}

func (c *Context) Bind(req interface{}) error {
	return c.GC.ShouldBind(req)
}

func (c *Context) Param(key string) string {
	return c.GC.Param(key)
}

/// Response

func (c *Context) ResponseError(err error) {
	c.GC.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  err.Error(),
		"data": "",
	})
}
func (c *Context) ResponseOK(data interface{}) {
	c.GC.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func (c *Context) ResponseList(data interface{}, page, count int64) {
	c.GC.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{
			"results": data,
			"page":    page,
			"count":   count,
		},
	})
}
