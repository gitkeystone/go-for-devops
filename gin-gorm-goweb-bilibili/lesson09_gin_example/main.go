package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 静态文件路由，确保在动态路由之前定义
	r.Static("/static", "./static")

	// 加载 HTML 模板文件
	r.LoadHTMLGlob("template/*.html")

	// 动态页面路由，改用更明确的路径，避免和静态文件路由冲突
	r.GET("/:page", func(c *gin.Context) {
		// 获取 URL 路径参数
		page := c.Param("page")
		if page == "" {
			page = "index" // 默认首页
		}

		// 渲染对应的模板
		c.HTML(http.StatusOK, page, nil)
	})

	// 启动服务器
	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("Server start failed, err: %v\n", err)
		return
	}
}
