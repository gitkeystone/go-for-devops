package main

// querystring

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET请求 URL ?后面是 querystring 参数
	// key=value格式，多个key-value用 & 链接
	// e.g.: /web/?name=小王子&age=18
	r.GET("/web", func(c *gin.Context) {
		//获取 querystring
		name := c.Query("name")
		age := c.Query("age")
		//name := c.DefaultQuery("name", "somebody")
		//name, ok := c.GetQuery("name")
		//if !ok {
		//	// 取不到
		//	name = "someone"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
