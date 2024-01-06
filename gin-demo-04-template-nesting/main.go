package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	k := func(name string) (string, error) {
		return name + "年轻又帅气！", nil
	}
	// 定义模板
	t := template.New("f.tmpl") // 创建一个名字是f的模板对象，名字一定要与模板的名字能对应上
	// 告诉模板引擎，我现在多了一个自定义的函数kua99
	t.Funcs(template.FuncMap{
		"kua99": k,
	})
	// 解析模板
	_, err := t.ParseFiles("./gin-demo-04-template-nesting/f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "程序员千羽"
	// 渲染模板
	t.Execute(w, name)

}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./gin-demo-04-template-nesting/t.tmpl", "./gin-demo-04-template-nesting/ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "程序员千羽"
	// 渲染模板
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/demo1", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
