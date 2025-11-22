package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Position struct {
	X int
	Y int
}

func distance(p1, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) +
		math.Pow(float64(p1.Y-p2.Y), 2))
}

type Player struct {
	ID   int
	Name string
	Position
}

func (p *Player) ShowPosition() {
	fmt.Printf("玩家：%s 的位置在 (%d, %d)\n", p.Name, p.X, p.Y)
}

type Monster struct {
	ID    int
	Speed int
	Position
}

func (m *Monster) moveTowards(playerPos Position) {
	dx := playerPos.X - m.X
	dy := playerPos.Y - m.Y
	dist := distance(m.Position, playerPos)

	if dist > 0 {
		m.Position.X += int(float64(dx) / dist * float64(m.Speed))
		m.Position.Y += int(float64(dy) / dist * float64(m.Speed))
	}

	fmt.Printf("怪物 %d 移动到: (%d, %d)\n", m.ID, m.X, m.Y)
}

func main() {
	player := Player{
		ID:       1,
		Name:     "Keystone",
		Position: Position{15, 10},
	}

	monster := Monster{
		ID:       1,
		Position: Position{X: 5, Y: 5},
		Speed:    2,
	}

	player.ShowPosition()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			monster.moveTowards(player.Position)
			if player.Position.X == monster.Position.X &&
				player.Position.Y == monster.Position.Y {
				fmt.Printf("怪物 1 遇上了玩家：%s，准备攻击...\n", player.Name)
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
