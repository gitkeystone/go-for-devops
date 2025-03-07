package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"lesson25_26_27_bubble/model"
)

var (
	DB  = new(gorm.DB)
	DSN = "todo.db"
)

func InitSqlite() (err error) {
	DB, err = gorm.Open(sqlite.Open(DSN), &gorm.Config{})
	return
}

func InitModel() error {
	return DB.AutoMigrate(&model.Todo{}) // todos
}

/*
	Todo CRUD
*/

// AddTodo 添加todo
func AddTodo(todo *model.Todo) error {
	return DB.Create(todo).Error
}

func GetATdo(todo *model.Todo, id string) error {
	return DB.Where("id = ?", id).First(todo).Error
}

func GetAllTodo(todos *[]model.Todo) error {
	return DB.Find(todos).Error
}

func UpdateATodo(todo *model.Todo) error {
	return DB.Save(todo).Error
}

func DeleteATodo(id string) error {
	return DB.Where("id = ?", id).Delete(&model.Todo{}).Error
}
