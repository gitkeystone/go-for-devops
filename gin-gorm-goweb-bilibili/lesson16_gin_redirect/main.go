package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK!",
		})
	})

	// 跳转到 网站
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.sogou.com/")
	})

	r.GET("/a", func(c *gin.Context) {
		// 跳转到 /b 对应的路径处理函数
		c.Request.URL.Path = "/b" // 修改请求的 URI
		r.HandleContext(c)        // 继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
