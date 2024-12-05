package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// UserInfo -> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open(sqlite.Open("db1"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移，创建表，就是把结构体和数据表对应
	_ = db.AutoMigrate(&UserInfo{})

	//u1 := UserInfo{1, "q1mi", "男", "蛙泳"}
	//db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("%#v\n", u)

	// 更新
	//db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)

}
