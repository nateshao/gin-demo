package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/hello", sayHello)
	r.Run()
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Golang!!",
	})

}
