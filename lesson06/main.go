package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", f1)
	err := http.ListenAndServe("8081", nil)
	if err != nil {
		fmt.Printf("parse template fail,err:%v", err)
		return
	}

}

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	//解析模板

	//渲染模板

}
