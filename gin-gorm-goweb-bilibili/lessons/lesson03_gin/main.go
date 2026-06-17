package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() // 返回默认的路由引擎

	//指定用户使用 GET 请求访问 /hello, 执行匿名函数
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Hello Golang!",
	//	})
	//})

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	//启动服务
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
