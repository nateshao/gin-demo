package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	gencer string
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("8081", nil)
	if err != nil {
		fmt.Println("HTTP server start fail ,err%v\n", err)
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template fail,err:%v", err)
		return
	}

	u := User{
		Name:   "张三",
		Age:    18,
		gencer: "男",
	}

	// 渲染模板
	t.Execute(w, u)

}
