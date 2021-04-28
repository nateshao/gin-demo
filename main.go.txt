//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	b, _ := ioutil.ReadFile("./hello.txt")
//	fmt.Fprintln(w, b)
//}
//
//func main() {
//	http.HandleFunc("/hello", sayHello)
//	err := http.ListenAndServe(":8085", nil)
//	if err != nil {
//		fmt.Printf("http fail..err:%v\n", err)
//		return
//	}
//
//}
/****************************************** gin示例 ******************************************************/
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
