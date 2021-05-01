package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//// HTTP 重定向很容易。 内部、外部重定向均支持。
	//r.GET("/test", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	//})

	// 路由重定向，使用HandleContext：
	r.GET("/test", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	r.Run()
}
