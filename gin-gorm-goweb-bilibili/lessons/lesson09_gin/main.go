package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 静态文件：
// html 页面上用到的样式文件 .css、js、图片
func main() {
	r := gin.Default()
	//在解析模板之前
	// 1. 加载静态文件
	r.Static("/xxx", "static")
	// 2. gin 框架中添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	//解析模板
	//r.LoadHTMLFiles("templates/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")

	//路由
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{ //渲染模板
			"title": "liwenzhou",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{ //渲染模板
			//"title": "users/index.tmpl",
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("Start Server failed, err: %v\n", err)
		return
	}
}
