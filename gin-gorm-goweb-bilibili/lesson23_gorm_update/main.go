package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 定义模型
type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	//连接 sqlite 数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//把模型与数据库中的表对应起来
	_ = db.AutoMigrate(&User{})

	//创建
	//users := []*User{
	//	{Name: "q1mi", Age: 18, Active: true},
	//	{Name: "jinzhu", Age: 20, Active: false},
	//}
	//db.Create(users)

	//查询
	var user User
	db.Debug().First(&user) // 查询第一条

	//更新
	//user.Name = "七米" // 仅对上述查询的记录进行修改
	//user.Age = 99
	//db.Debug().Save(&user)                        // 默认修改所有字段
	//db.Debug().Model(&user).Update("name", "小王子") // 仅修改指定字段
	//
	//m1 := map[string]any{
	//	"name":   "chenxiaohui",
	//	"age":    28,
	//	"active": true,
	//}
	//db.Debug().Model(&user).Updates(m1)
	//db.Debug().Model(&user).Select("age").Updates(m1)  // 仅更新age字段
	//db.Debug().Model(&user).Omit("active").Updates(m1) // 排除 active 字段
	//
	//db.Debug().Model(&user).UpdateColumn("age", 30) // 不触发 Hooks
	//affected := db.Debug().Model(User{}).Where("name = ?", "jinzhu").Updates(User{Name: "hello", Age: 18}).RowsAffected
	//fmt.Println(affected)

	//让 users 所有用户的年龄+2
	db.Debug().Model(&user).Update("age", gorm.Expr("age + ?", 2))

}
