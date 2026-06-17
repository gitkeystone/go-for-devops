package main

import (
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 定义模型
type User struct {
	ID int64
	//Name string `gorm:"default:'小王子'"`
	//Name *string `gorm:"default:'小王子'"`
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&User{})

	// 创建记录
	//u := User{Name: "q1mi", Age: 18}
	//u := User{Age: 38} // 使用默认值
	//u := User{Name: new(string), Age: 38} // 零值 == 不写；
	//u := User{Name: sql.NullString{Valid: true}, Age: 38} // 零值 == 不写；
	u := User{Name: sql.NullString{Valid: true}, Age: 0} // 零值 == 不写；

	//db.Create(&u)
	db.Debug().Create(&u)
}
