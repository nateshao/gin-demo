package main

import (
	"fmt"
	"io/ioutil"
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
	file, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintf(w, string(file))
	fmt.Println(w, "<h1>Hello Golang</h1>")
}
