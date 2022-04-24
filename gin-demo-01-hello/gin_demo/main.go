package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}

func hello(h http.ResponseWriter, r *http.Request) {
	// 读取文件
	b, _ := ioutil.ReadFile("./hello.txt")
	//打印
	fmt.Fprintln(h, string(b))
}

func main() {
	//r := gin.Default() // 返回默认的路有引擎
	//
	//// 指定用户使用GET请求访问/hello时，执行sayHello这个函数
	//r.GET("/sayHello", sayHello)
	//
	//// 启动服务
	//r.Run(":8080")

	/************************************/

	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failed ,err :%v\n", err)
		return
	}

	//r.GET("/book", ...)
	//r.GET("/create_book", ...)
	//r.GET("/update_book", ...)
	//r.GET("/shanchu_book", ...)

	//r.GET("/book", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"method": "GET",
	//	})
	//})
	//r.POST("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "POST",
	//	})
	//})
	//r.PUT("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "PUT",
	//	})
	//})
	//r.DELETE("/book", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "DELETE",
	//	})
	//})
}
