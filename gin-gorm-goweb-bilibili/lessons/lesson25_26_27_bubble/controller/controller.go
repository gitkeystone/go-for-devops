package controller

import (
	"github.com/gin-gonic/gin"
	"lesson25_26_27_bubble/dao"
	"lesson25_26_27_bubble/model"
	"net/http"
)

/*
	URL -> Controller -> logic -> model
	请求来了 -》控制器 -》响应层的业务逻辑-》模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	//前端填写代办事项，点击提交
	//1.从请求中，提取数据
	var todo model.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}

	//2.存入数据库
	err := dao.AddTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg":  "success",
		//	"data": todo,
		//})
	}

}

func GetTodoList(c *gin.Context) {
	// 查询 todos 表中所有数据
	var todos []model.Todo
	err := dao.GetAllTodo(&todos)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func UpdateATodo(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var todo model.Todo
	err := dao.GetATdo(&todo, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的 id",
		})
		return
	}

	_ = c.BindJSON(&todo)

	err = dao.UpdateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的 ID",
		})
		return
	}

	err := dao.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}
