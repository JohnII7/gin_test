package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		fmt.Println("parse template error:", err)
		return
	}

	// 渲染模板
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    24,
	}
	m1 := map[string]interface{}{
		"Name":   "老王子",
		"Gender": "女",
		"Age":    17,
	}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server start error:", err)
		return
	}
}
