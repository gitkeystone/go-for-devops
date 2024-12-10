package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	ID     int
	Name   string
	Age    int64
	Active bool
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(User{})

	//users := []*User{
	//	{Name: "q1mi", Age: 18, Active: true},
	//	{Name: "jinzhu", Age: 20, Active: false},
	//}
	//db.Create(users)

	//删除
	var u User
	u.ID = 1
	//u.Name = "jinzhu"
	db.Debug().Delete(&u)

	//db.Debug().Where("name = ?", "jinzhu").Delete(&User{})
	//db.Debug().Delete(&User{}, "name = ?", 18)

	// 查询软删除
	//var u1 []User
	//db.Where("name = ?", "q1mi").Find(&u1)
	//db.Unscoped().Where("name = ?", "q1mi").Find(&u1)
	//fmt.Println(u1)

	//物理删除
	//db.Debug().Unscoped().Where("name = ?", "q1mi").Delete(&User{})
}
