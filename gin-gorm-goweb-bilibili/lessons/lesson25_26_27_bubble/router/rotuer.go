package router

import (
	"github.com/gin-gonic/gin"
	"lesson25_26_27_bubble/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")

	r.GET("/", controller.IndexHandler)

	// 待办事项
	// v1
	v1Group := r.Group("v1")
	{
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看
		v1Group.GET("/todo", controller.GetTodoList)

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
