package main

import "fmt"

type Equipment struct {
	Name   string
	Power  int
	Defend int
}

type Character struct {
	Name      string
	Level     int
	Health    int
	Equipment *Equipment
}

func LevelUp(character *Character) {
	character.Level++
	fmt.Printf("%s 升级到了 %d 级\n", character.Name, character.Level)
}

func ChangeEquipment(character *Character, newEquipment *Equipment) {
	character.Equipment = newEquipment
	fmt.Printf("%s 更换了装备为 %s\n", character.Name, newEquipment.Name)
}

func main() {
	sword := Equipment{
		Name:   "铁剑",
		Power:  10,
		Defend: 2,
	}

	warrior := Character{
		Name:      "赵云",
		Level:     1,
		Health:    100,
		Equipment: &sword,
	}

	LevelUp(&warrior)

	newSword := Equipment{
		Name:   "青铜剑",
		Power:  20,
		Defend: 5,
	}

	ChangeEquipment(&warrior, &newSword)

}
