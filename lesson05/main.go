package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("9000", nil)
	if err != nil {
		fmt.Println("HTTP server start fail ,err", err)

	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template fail,err:%v", err)
		return
	}

	// 渲染模板
	t.Execute(w, "是千羽哦")

}
