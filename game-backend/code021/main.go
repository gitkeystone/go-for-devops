package main

import "fmt"

type Character struct {
	Name   string
	HP     int
	Attack int
	Defend int
}

func (c *Character) AttackTarget(target *Character) {
	damage := c.Attack - target.Defend
	if damage < 0 {
		damage = 0
	}
	target.HP -= damage
	fmt.Printf("%s 攻击了 %s，造成了 %d 点伤害，%s 剩余HP: %d\n", c.Name, target.Name, damage, target.Name, target.HP)
}

func Battle(player, monster *Character) {
	for player.HP > 0 && monster.HP > 0 {
		player.AttackTarget(monster)
		if monster.HP > 0 {
			monster.AttackTarget(player)
		}
	}

	if monster.HP <= 0 {
		fmt.Println(player.Name, "战胜了", monster.Name)
	}

	if player.HP <= 0 {
		fmt.Println(monster.Name, "战胜了", player.Name)
	}
}

func main() {
	player := Character{
		Name:   "勇者",
		HP:     100,
		Attack: 20,
		Defend: 10,
	}

	monster := Character{
		Name:   "哥布林",
		HP:     80,
		Attack: 15,
		Defend: 5,
	}

	Battle(&player, &monster)
}
