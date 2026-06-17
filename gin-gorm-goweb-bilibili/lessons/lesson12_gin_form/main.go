package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostForm 获取表单提交的参数
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html", "index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// /login POST
	r.POST("/login", func(c *gin.Context) {
		// 获取 Form 表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")  // 取不到，返回类型默认值
		//username := c.DefaultPostForm("username", "somebody")	// 取不到，返回指定值-somebody
		//password := c.DefaultPostForm("xxx", "***")
		username, ok := c.GetPostForm("username") // 取不到，返回 false
		if !ok {
			username = "sb"
		}
		password, _ := c.GetPostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})

	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
