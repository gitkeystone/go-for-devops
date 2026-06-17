package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var r = gin.Default()
	//method1: map
	r.GET("/json", func(c *gin.Context) {
		//data := map[string]any{
		//	"name":    "小王子",
		//	"message": "hello world",
		//	"age":     18,
		//}

		data := gin.H{"name": "小王子", "message": "hello world", "age": 18}

		c.JSON(http.StatusOK, data)
	})
	// method2: struct
	type msg struct {
		Name    string `json:"name,omitempty"` // 首字母大写-才能被看到；打标签-使得首字母小写
		Message string `json:"message,omitempty"`
		Age     int    `json:"age,omitempty"`
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "小王子",
			Message: "hello world",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})
	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("HTTP start failed, err: %v\n", err)
		return
	}
}
