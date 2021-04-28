package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server fail..,err%v", err)
		return
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("http server fail..,err%v", err)
		return
	}
	name := "千羽的编程时光"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("template....fail..")
	}
	// 渲染模板
}
