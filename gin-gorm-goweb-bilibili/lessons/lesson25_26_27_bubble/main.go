package main

import (
	"lesson25_26_27_bubble/dao"
	"lesson25_26_27_bubble/router"
)

func main() {
	//连接数据库
	err := dao.InitSqlite()
	if err != nil {
		panic(err)
	}

	err = dao.InitModel()
	if err != nil {
		panic(err)
	}

	r := router.SetupRouter()
	err = r.Run(":9000")
	if err != nil {
		panic("failed to run HTTP server")
	}
}
