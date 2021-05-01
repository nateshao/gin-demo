package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Genter string
}

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
	u := User{
		Name:   "千羽",
		Age:    18,
		Genter: "男",
	}
	//u := map[string]interface{}{
	//	Name:   "千羽",
	//	Age:    18,
	//	Genter: "男",
	//}
	//name := "千羽的编程时光"
	// 渲染模板
	err = t.Execute(w, u)
	if err != nil {
		fmt.Println("template....fail..")
	}
}
