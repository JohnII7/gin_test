package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取请求携带的queryString参数
		name := c.Query("query")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("err:", err)
	}
}
