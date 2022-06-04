package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "邵桐杰"
	// 渲染模板
	t.Execute(w, msg)
}

type User struct {
	Name   string
	Gender string
	Age    int
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	u1 := User{
		Name:   "千羽",
		Gender: "男",
		Age:    18,
	}

	m1 := map[string]interface{}{
		"name":   "邵桐杰",
		"gender": "男",
		"age":    18,
	}

	//msg := "邵桐杰"
	// 渲染模板
	//t.Execute(w, msg)
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}

	msg := "邵桐杰"
	// 渲染模板
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板（模板继承的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	// 渲染模板
	name := "邵桐杰"
	t.ExecuteTemplate(w, "index2.tmpl", name)
}
func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板（模板继承的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	// 渲染模板
	name := "千羽"
	t.ExecuteTemplate(w, "home2.tmpl", name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
