package main

import "fmt"

func main() {
	num1 := 10
	num2 := 3
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2
	remainder := num1 % num2

	fmt.Printf("和：%d\n", sum)
	fmt.Printf("差：%d\n", difference)
	fmt.Printf("积：%d\n", product)
	fmt.Printf("商：%d\n", quotient)
	fmt.Printf("余数：%d\n", remainder)

	level1 := 10
	level2 := 15
	isEqual := level1 == level2
	isNotEqual := level1 != level2
	isGreater := level1 > level2
	isLess := level1 < level2
	fmt.Printf("level1 等于 level2: %v\n", isEqual)
	fmt.Printf("level1 不等于 level2: %v\n", isNotEqual)
	fmt.Printf("level1 大于 level2: %v\n", isGreater)
	fmt.Printf("level1 小于 level2: %v\n", isLess)

	condition1 := true
	condition2 := false
	andResult := condition1 && condition2
	orResult := condition1 || condition2
	notResult := !condition1
	fmt.Printf("逻辑与结果：%v\n", andResult)
	fmt.Printf("逻辑或结果：%v\n", orResult)
	fmt.Printf("逻辑非结果：%v\n", notResult)
}
