package main

import "fmt"

// Weapon 定义武器结构体
type Weapon struct {
	Name        string
	Description string
	Damage      int
}

// Armor 定义防具结构体
type Armor struct {
	Name        string
	Description string
	Defense     int
}

// Item 定义物品接口
type Item interface {
	GetName() string
	GetDescription() string
}

func (w Weapon) GetName() string {
	return w.Name
}

func (w Weapon) GetDescription() string {
	return fmt.Sprintf("%s - 攻击力：%d", w.Description, w.Damage)
}

func (a Armor) GetName() string {
	return a.Name
}

func (a Armor) GetDescription() string {
	return fmt.Sprintf("%s - 防御力: %d", a.Description, a.Defense)
}

// PlayerInventory 玩家背包结构体
type PlayerInventory struct {
	Items []Item
}

// AddItem 向背包添加物品
func (pi *PlayerInventory) AddItem(item Item) {
	pi.Items = append(pi.Items, item)
}

// ShowInventory 展示背包物品
func (pi *PlayerInventory) ShowInventory() {
	fmt.Println("背包物品:")
	for _, item := range pi.Items {
		fmt.Printf("- %s: %s\n", item.GetName(), item.GetDescription())
	}
}

func main() {
	playerInventory := PlayerInventory{}
	// 创建武器和防具实例
	sword := Weapon{
		Name:        "铁剑",
		Description: "一把普通的铁剑",
		Damage:      20,
	}
	leatherArmor := Armor{
		Name:        "皮甲",
		Description: "基础的皮制防具",
		Defense:     10,
	}

	playerInventory.AddItem(sword)
	playerInventory.AddItem(leatherArmor)

	playerInventory.ShowInventory()
}
