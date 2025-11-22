package main

import "fmt"

type Task struct {
	Name        string
	Description string
}

type Item struct {
	Name string
	Task *Task
}

type Character struct {
	Name      string
	Inventory []*Item
}

func PickUpItem(character *Character, item *Item) {
	character.Inventory = append(character.Inventory, item)
	fmt.Printf("%s 捡起了 %s\n", character.Name, item.Name)
}

func ShowTaskItems(character *Character) {
	for _, item := range character.Inventory {
		if item.Task != nil {
			fmt.Printf("%s 拥有与任务 %s 相关的物品 %s\n", character.Name, item.Task.Name, item.Name)
		}
	}
}

func main() {
	mainTask := Task{
		Name:        "寻找神秘宝石",
		Description: "在古老遗迹中寻找神秘宝石",
	}

	gem := Item{
		Name: "神秘宝石",
		Task: &mainTask,
	}

	player := Character{
		Name:      "冒险者",
		Inventory: make([]*Item, 0),
	}

	PickUpItem(&player, &gem)

	ShowTaskItems(&player)
}
