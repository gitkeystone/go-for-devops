package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求的 path (URI) 参数，返回的都是字符串类型
// 注意 URL 匹配不要冲突
func main() {
	r := gin.Default()

	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
