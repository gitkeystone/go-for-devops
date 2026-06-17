package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 定义模型
type User struct {
	gorm.Model
	Name string
	Age  uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&User{})

	// 创建两条记录
	//db.Create([]*User{
	//	{Name: "q1mi", Age: 18},
	//	{Name: "jinzhu", Age: 20},
	//})

	// 查询
	//var user User      // 声明
	//db.First(&user) // 按主键排序后，第一条记录
	//db.First(&user, 2) // 按主键排序后，第一条记录
	//fmt.Printf("%#v\n", user)

	//var users []User
	////db.Find(&users)
	//db.Debug().Find(&users)
	//fmt.Printf("%#v\n", users)

	// 如果没有查询到匹配记录，则初始化给定实例数据
	//var user User
	//db.Attrs(User{Age: 99}).FirstOrInit(&user, User{Name: "non-existing"})
	//db.Assign(User{Age: 20}).FirstOrInit(&user, User{Name: "jinzhu"})
	//fmt.Printf("%#v\n", user)

}
