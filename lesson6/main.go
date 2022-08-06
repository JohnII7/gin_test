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

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua
	kua := func(name string) string {
		return name + "真帅"
	}

	// 定义模板
	t, err := template.New("hello.html").ParseFiles("./hello.html")

	// 通知模板引擎多了一个自定义函数kua
	t.Funcs(template.FuncMap{
		"kua99": kua,
	})
	// 解析模板
	// 模板对象要和模板名字对应上
	if err != nil {
		fmt.Printf("parse error failed, error:%v\n", err)
		return
	}
	name := "小王子"
	// 渲染模板
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server start error:", err)
		return
	}
}
