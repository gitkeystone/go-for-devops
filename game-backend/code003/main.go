package main

import "fmt"

func main() {
	isWin := true
	fmt.Println("玩家是否胜利：", isWin)

	speed := 3.1415926 // float64
	fmt.Println("角色移动速度：", speed)

	hpArray := [3]int{100, 200, 300}
	fmt.Println("怪物生命数值组：", hpArray)

	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	fmt.Println("切片", numbers)
}
