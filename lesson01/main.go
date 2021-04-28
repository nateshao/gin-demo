package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("http server fail..,err%v", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "<h1>Hello Golang</h1>")
}
