package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")
	fmt.Fprintln(w, b)
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Printf("http fail..err:%v\n", err)
		return
	}

}
