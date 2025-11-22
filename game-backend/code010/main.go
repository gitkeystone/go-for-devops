package main

import "fmt"

// 定义一个攻击行为的接口
type Attacker interface {
	Attack() int
}

// 定义战士结构体，并实现 Attacker 接口
type Warrior struct {
	Name        string
	AttackPower int
}

func (w Warrior) Attack() int {
	return w.AttackPower
}

type Mage struct {
	Name       string
	MagicPower int
}

func (m Mage) Attack() int {
	return m.MagicPower
}

func doAttack(a Attacker) {
	damage := a.Attack()
	fmt.Printf("造成了%d点伤害\n", damage)
}

func main() {
	warrior := Warrior{
		Name:        "赵云",
		AttackPower: 50,
	}

	mage := Mage{
		Name:       "貂蝉",
		MagicPower: 30,
	}

	fmt.Printf("%s开始攻击：", warrior.Name)
	doAttack(warrior)
	fmt.Printf("%s开始攻击：", mage.Name)
	doAttack(mage)
}
