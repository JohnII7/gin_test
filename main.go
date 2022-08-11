package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法一 map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     24,
		//}

		data := gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     24,
		}
		c.JSON(http.StatusOK, data)
	})
	// 方法二 struct
	type msg struct {
		Name    string `json:"name"`
		Age     int
		Message string
	}

	r.GET("/anotherJson", func(context *gin.Context) {
		data := msg{
			Name:    "大王子",
			Age:     19,
			Message: "HelloWorld",
		}
		context.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
