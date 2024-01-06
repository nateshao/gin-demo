package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 遇事不决 写注释

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl") // 请勿刻舟求剑
	if err != nil {
		fmt.Println("Parse template failed,err:%v", err)
		return
	}
	// 3. 渲染模板
	name := "公众号：程序员千羽"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed,err:%v", err)
		return
	}

}

func sayNateshao(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./nateshao.tmpl")
	if err != nil {
		fmt.Println("解析文件失败:%v", err)
		return
	}
	name := "公众号千羽"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed,err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/nateshao", sayNateshao)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}

}
