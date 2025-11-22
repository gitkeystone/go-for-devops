package main

import "fmt"

// 定义一个表示游戏角色的结构体
type Character struct {
	Name   string
	Level  int
	HP     int
	Attack int
}

func heal(character *Character, amount int) {
	character.HP += amount
}

func main() {
	hero := Character{
		Name:   "战士keystone",
		Level:  10,
		HP:     100,
		Attack: 20,
	}
	//fmt.Printf("角色名称：%s，等级：%d，生命值：%d，攻击力：%d\n", hero.Name, hero.Level, hero.HP, hero.Attack)
	fmt.Printf("治疗前生命值：%d\n", hero.HP)
	heal(&hero, 50)
	fmt.Printf("治疗后生命值：%d\n", hero.HP)
}
