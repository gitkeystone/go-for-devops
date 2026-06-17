package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Printf("%s\n", "index in")
	name, ok := c.Get("name") // 取 key
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
	fmt.Println("index out")
}

// 定义中间件-统计耗时
func m1(c *gin.Context) {
	fmt.Printf("%s\n", "m1 in...")
	//计时
	start := time.Now()
	go func() {}() // 并发不安全；在 func 中只能使用 c 的拷贝 c.Copy()
	c.Next()       // 调用后续 handler
	//c.Abort() // 阻止后续 handler

	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Printf("%s\n", "m2 in")
	c.Set("name", "q1mi") // 设置 key-value
	//c.Next() // 调用后续 handler
	//c.Abort() // 阻止后续 handler
	//return    // 立即结束 m2
	println("m2 out")
}

// 标准的闭包格式
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//链接数据库
	//或其他一些准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体逻辑
			//是否登录成功
			//if 登录成功
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	//r := gin.Default() // 默认自带中间件 Logger/Recovery
	r := gin.New()

	//全局注册中间件函数 m1/m2/authMiddleware
	r.Use(m1, m2, authMiddleware(false))

	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	//路由组注册中间件方法1
	videoGroup := r.Group("/video", authMiddleware(true))
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "videoGroup",
			})
		})
	}
	//路由组注册中间件方法1
	webGroup := r.Group("/web")
	webGroup.Use(authMiddleware(true))
	{
		webGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "webGroup",
			})
		})
	}

	//运行
	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
