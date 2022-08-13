package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler method")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件m1, 统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 method")
	// 计时
	start := time.Now()
	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止后续的处理函数
	cost := time.Since(start)
	fmt.Println("cost time :", cost)
}
func m2(c *gin.Context) {
	fmt.Println("m2 method")
	c.Next() // 调用后续的处理函数
	fmt.Println("m2 method end")
}

func main() {
	r := gin.Default()
	// 全局注册中间件函数
	r.Use(m1, m2)
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop method",
		})
		fmt.Println("shop method")
	})
	r.GET("/video", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "video method",
		})
		fmt.Println("video method")
	})
	_ = r.Run()
}
