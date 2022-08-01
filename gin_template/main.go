package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("Parse err:", err)
		return
	}

	// 渲染模板
	err = t.Execute(w, "小王子")
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
	//
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server start error:", err)
		return
	}
}
