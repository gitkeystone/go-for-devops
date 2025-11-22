package main

import "fmt"

type Player struct {
	Name    string
	Balance int
	Items   []string
}

func (p *Player) EarnMoney(amount int) {
	p.Balance += amount
	fmt.Printf("%s 获得了 %d 金币，当前余额: %d\n", p.Name, amount, p.Balance)
}

func (p *Player) SpendMoney(amount int) bool {
	if p.Balance >= amount {
		p.Balance -= amount
		fmt.Printf("%s 花费了 %d 金币，当前余额: %d\n", p.Name, amount, p.Balance)
		return true
	}
	fmt.Printf("%s 余额不足，无法花费 %d 金币\n", p.Name, amount)
	return false
}

type Store struct {
	Items map[string]int
}

func (s *Store) SellItem(player *Player, itemName string, price int) {
	if _, ok := s.Items[itemName]; ok {
		if player.SpendMoney(price) {
			player.Items = append(player.Items, itemName)
			fmt.Printf("%s 从商店购买了 %s\n", player.Name, itemName)
		}
	} else {
		fmt.Printf("商店没有 %s 出售\n", itemName)
	}
}

func main() {
	player := Player{
		Name:    "玩家A",
		Balance: 100,
		Items:   []string{},
	}

	// 创建商店
	store := Store{
		Items: map[string]int{
			"剑":  50,
			"盾":  40,
			"药水": 20,
		},
	}

	player.EarnMoney(50)

	store.SellItem(&player, "剑", 50)
	store.SellItem(&player, "铠甲", 60)

	fmt.Printf("%s 的物品: %v\n", player.Name, player.Items)
}
