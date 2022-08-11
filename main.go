package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "John" && password == "123" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Name":     username,
				"Password": password,
			})
		}

	})

	err := r.Run(":9090")
	if err != nil {
		fmt.Println("err:", err)
	}
}
