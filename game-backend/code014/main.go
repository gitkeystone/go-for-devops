package main

import "fmt"

func main() {
	scores := make(map[string]int)
	scores["赵云"] = 92
	scores["不知火舞"] = 88
	fmt.Println(scores)

	cities := map[string]int{
		"房间1": 15,
		"房间2": 28,
	}
	fmt.Println(cities)

	fruits := map[string]int{
		"苹果": 5,
		"西瓜": 3,
	}
	count, ok := fruits["苹果"]
	if ok {
		fmt.Printf("map中有%d个苹果\n", count)

	} else {
		fmt.Println("该map中没有找到苹果")
	}
	count, ok = fruits["香蕉"]
	if ok {
		fmt.Printf("map中有%d根香蕉\n", count)
	} else {
		fmt.Println("该map中没有找到香蕉")
	}

	animals := map[string]int{
		"小狗": 2,
		"小猫": 3,
		"小鸟": 5,
	}
	fmt.Println("删除前：", animals)
	delete(animals, "小猫")
	fmt.Println("删除后：", animals)

	numbers := map[string]int{
		"一": 1,
		"二": 2,
		"三": 3,
	}
	for key, value := range numbers {
		fmt.Printf("%s: %d\n", key, value)
	}

}
