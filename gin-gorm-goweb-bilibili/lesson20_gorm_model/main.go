package main

import (
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	gorm.Model   // 内嵌
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` // 零值
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// Animal 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// TableName 自定义数据库中的表名
func (Animal) TableName() string {
	return "custom_animal"
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,   // 禁用默认复数表名
			TablePrefix:   "SMS_", // 添加默认表名前缀
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&User{})
	_ = db.AutoMigrate(&Animal{})

}
