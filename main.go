package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		// 请求中读取单文件
		form, _ := c.MultipartForm()
		files := form.File["f1"]
		for i, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, i)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	err := r.Run(":9090")
	if err != nil {
		fmt.Println("err:", err)
	}
}
