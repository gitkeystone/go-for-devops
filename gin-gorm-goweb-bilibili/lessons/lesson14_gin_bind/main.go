package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"user,omitempty"`
	Password string `form:"password" json:"pwd,omitempty"`
}

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}

		var u UserInfo          // 变量声明, 值类型
		err := c.ShouldBind(&u) // 传递引用类型 &u, 不能是 u
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}

	})

	r.LoadHTMLFiles("index.html")
	r.GET("index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo          // 变量声明, 值类型
		err := c.ShouldBind(&u) // 传递引用类型 &u, 不能是 u
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u UserInfo          // 变量声明, 值类型
		err := c.ShouldBind(&u) // 传递引用类型 &u, 不能是 u
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
