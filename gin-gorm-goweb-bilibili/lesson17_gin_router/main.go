package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 普通路由:GET/POST/PUT/DELETE/Any/NoRoute
	// 什么是路由？访问 /index 的 GET 请求，会走这一条处理逻辑-函数

	//查询
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodGet,
		})
	})
	//创建
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodPost,
		})
	})

	//修改
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodPut,
		})
	})

	//删除
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": http.MethodDelete,
		})
	})

	//所有
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": http.MethodGet})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": http.MethodPost})
		}
	})

	//无路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"method": "NoRoute",
		})
	})

	//路由组
	//视频的首页、详情页
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}

	//商城的首页、详情页
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
		})
		shopGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/xx"})
		})
		shopGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/oo"})
		})
	}

	//运行
	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
